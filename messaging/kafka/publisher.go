package kafka

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/super-saga/go-x/graceful"
	"github.com/super-saga/go-x/messaging"

	"github.com/IBM/sarama"
)

type publisher struct {
	producer sarama.AsyncProducer
	origin   string
	wg       sync.WaitGroup
	metrics  *publisherPrometheusMetrics
}

// Create new instance of *publisher.
// This struct is implement several interface in messaging package.
// You may specify the opts by passing messaging.WithOrigin("yourservicename")
// to automatically add origin in your header.
func NewPublisher(brokers []string, opts ...PublisherOption) (pub *publisher, stopper graceful.ProcessStopper, err error) {
	stopper = func(context.Context) error { return nil }
	pub = new(publisher)
	opts = append(defaultPublisherOptions, opts...)
	popt := newPublisherOption()
	popt.applyOptions(opts...)

	producer, err := sarama.NewAsyncProducer(brokers, popt.saramaCfg)
	if err != nil {
		err = fmt.Errorf("error creating sarama producer: %w", err)
		return
	}

	pub.producer = producer
	pub.origin = popt.origin
	pub.metrics = popt.metrics

	// Run default kafka metrics collector
	if pub.metrics != nil {
		pub.metrics.Run()
	}

	pub.run()

	stopper = func(ctx context.Context) error {
		return pub.ClosePublisher()
	}

	return
}

// PublisherCloser implementation.
func (p *publisher) ClosePublisher() error {
	p.producer.AsyncClose()
	p.wg.Wait()
	return nil
}

// PublisherTransaction implementation.
func (p *publisher) IsTransactional() bool {
	return p.producer.IsTransactional()
}

// PublisherTransaction implementation.
func (p *publisher) BeginTxn() error {
	return p.producer.BeginTxn()
}

// PublisherTransaction implementation.
func (p *publisher) CommitTxn() error {
	return p.producer.CommitTxn()
}

// PublisherTransaction implementation.
func (p *publisher) AbortTxn() error {
	return p.producer.AbortTxn()
}

// PublisherTransaction implementation.
func (p *publisher) StatusTxn() messaging.PublisherTxnStatusFlag {
	return messaging.PublisherTxnStatusFlag(p.producer.TxnStatus())
}

// PublisherTransaction implementation.
func (p *publisher) AddMessageToTxn(msg messaging.Message, groupID string) error {
	consumerMsg, ok := msg.(*Message)
	if !ok {
		return fmt.Errorf("error assert type to *kafka.Message")
	}
	return p.producer.AddMessageToTxn(consumerMsg.ConsumerMessage, groupID, nil)
}

func (p *publisher) PublishSyncWithKey(ctx context.Context, topic string, key string, body interface{}, opts ...messaging.PublishOption) error {
	return p.PublishSyncWithKeyAndHeader(ctx, topic, key, body, nil, opts...)
}

func (p *publisher) PublishSyncWithKeyAndHeader(ctx context.Context, topic string, key string, body interface{}, headers map[string]interface{}, opts ...messaging.PublishOption) error {
	promise, err := p.PublishAsyncWithKeyAndHeader(ctx, topic, key, body, headers, opts...)
	if err != nil {
		return err
	}

	done := make(chan struct{})
	promise.Then(func(asyncErr error) {
		if asyncErr != nil {
			err = fmt.Errorf("error publish message: %w", asyncErr)
		}
		close(done)
	})
	<-done

	return err
}

func (p *publisher) PublishAsyncWithKey(ctx context.Context, topic string, key string, body interface{}, opts ...messaging.PublishOption) (*messaging.Promise, error) {
	return p.PublishAsyncWithKeyAndHeader(ctx, topic, key, body, nil, opts...)
}

func (p *publisher) PublishAsyncWithKeyAndHeader(ctx context.Context, topic string, key string, body interface{}, headers map[string]interface{}, opts ...messaging.PublishOption) (*messaging.Promise, error) {
	var (
		startTime = time.Now()
	)

	// ctx, dt := newrelic.BeforePublish(ctx, fmt.Sprintf(PublisherSegmentPrefixFmt, topic))
	// defer newrelic.AfterPublish(ctx)
	// newrelic.AddSegmentAttrs(ctx, newrelic.WithSegmentAttr(AttrPartitionKey, key))

	opts = append([]messaging.PublishOption{messaging.WithDefaultCodec()}, opts...)
	opt := messaging.NewPublishOption()
	opt.ApplyOption(opts...)

	bodyByte, err := opt.Codec().Encode(body)
	if err != nil {
		return nil, fmt.Errorf("error marshall body: %w", err)
	}

	recordHeaders := []sarama.RecordHeader{
		{
			Key:   []byte("content-type"),
			Value: []byte(opt.Codec().ContentType()),
		},
		{
			Key:   []byte("schema-version"),
			Value: []byte(opt.Codec().SchemaVersion()),
		},
	}
	if len(p.origin) > 0 {
		recordHeaders = append(recordHeaders, sarama.RecordHeader{
			Key:   []byte("origin"),
			Value: []byte(p.origin),
		})
	}
	// for key, val := range dt.Header {
	// 	recordHeaders = append(recordHeaders, sarama.RecordHeader{
	// 		Key:   []byte(key),
	// 		Value: []byte(val[0]),
	// 	})
	// }
	for key, val := range headers {
		if val == nil {
			continue
		}
		valByte, err := opt.Codec().Encode(val)
		if err != nil {
			return nil, fmt.Errorf("error marshall header value: %w", err)
		}
		recordHeaders = append(recordHeaders, sarama.RecordHeader{
			Key:   []byte(key),
			Value: valByte,
		})
	}

	promise := messaging.NewPromise()
	p.producer.Input() <- &sarama.ProducerMessage{
		Topic:    topic,
		Key:      sarama.StringEncoder(key),
		Value:    sarama.ByteEncoder(bodyByte),
		Headers:  recordHeaders,
		Metadata: promise,
	}

	if p.metrics != nil {
		promise = promise.Then(func(err error) {
			p.metrics.Record(topic, time.Since(startTime))
		})
	}

	return promise, nil
}

func (p *publisher) run() {
	p.wg.Add(2)
	go func() {
		defer p.wg.Done()
		for err := range p.producer.Errors() {
			err.Msg.Metadata.(*messaging.Promise).Finish(err)
		}
	}()
	go func() {
		defer p.wg.Done()
		for msg := range p.producer.Successes() {
			msg.Metadata.(*messaging.Promise).Finish(nil)
		}
	}()
}
