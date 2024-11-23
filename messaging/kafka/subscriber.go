package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/super-saga/go-x/messaging"
	"golang.org/x/sync/errgroup"
)

type subscriber struct {
	consumerGroup sarama.ConsumerGroup
	logger        sarama.StdLogger
	handlers      map[string]consumerHandler
	mu            sync.RWMutex
	middlewares   []messaging.MiddlewareFunc
	errCb         func(error)
	preStartCb    func(context.Context, map[string][]int32) error
	endCb         func(context.Context, map[string][]int32) error
	metrics       *subscriberPrometheusMetrics
}

type subscriberHandler struct {
	*subscriber
}

func NewSubscriber(brokers []string, groupID string, opts ...SubscriberOption) (*subscriber, error) {
	opts = append(defaultSubscriberOptions, opts...)
	opt := newSubscriberOption()
	opt.applyOptions(opts...)

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, opt.saramaCfg)
	if err != nil {
		return nil, fmt.Errorf("error creating new consumer group: %w", err)
	}

	sub := &subscriber{
		consumerGroup: consumerGroup,
		logger:        opt.saramaLog,
		handlers:      make(map[string]consumerHandler),
		middlewares:   opt.middlewares,
		errCb:         opt.errCb,
		preStartCb:    opt.preStartCb,
		endCb:         opt.endCb,
		metrics:       opt.metrics,
	}

	// Run default kafka metrics collector
	if opt.metrics != nil {
		opt.metrics.Run()
		sub.middlewares = append(sub.middlewares, opt.metrics.Middleware(groupID))
	}

	return sub, nil
}

func (s *subscriber) Subscribe(ctx context.Context, topics ...messaging.Topic) error {
	if s.errCb != nil {
		go s.handleError()
	}
	sh := &subscriberHandler{
		subscriber: s,
	}

	topicNames := s.populateTopics(topics)
	errg, ctx := errgroup.WithContext(ctx)

	errg.Go(func() error {
		for {
			if err := s.consumerGroup.Consume(ctx, topicNames, sh); err != nil {
				s.logger.Printf("consumer group consume error: %s", err)
			}
			if err := ctx.Err(); err != nil {
				err = fmt.Errorf("context was canceled: %s", err)
				s.logger.Printf(err.Error())
				return err
			}
		}
	})

	return errg.Wait()
}

func (s *subscriber) CloseSubscriber() error {
	for _, h := range s.handlers {
		h.Close()
	}
	return s.consumerGroup.Close()
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (s *subscriberHandler) Setup(session sarama.ConsumerGroupSession) error {
	for _, h := range s.handlers {
		h.Begin(session)
	}
	return s.preStartCb(session.Context(), session.Claims())
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (s *subscriberHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	for _, h := range s.handlers {
		h.End(session)
	}
	return s.endCb(session.Context(), session.Claims())
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (s *subscriberHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	s.mu.Lock()
	handler, handlerExists := s.handlers[claim.Topic()]
	s.mu.Unlock()
	if !handlerExists {
		return fmt.Errorf("error handler not exists: topic %s", claim.Topic())
	}
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			handler.Handle(session, message, func(msg messaging.Message, topic messaging.Topic) (err error) {
				defer func() {
					if rec := recover(); rec != nil {
						err = fmt.Errorf("panic detected! %v", err)
						s.sendErrorCb(err)
					}
				}()
				middlewares := append(s.middlewares, topic.Middlewares()...)
				h := s.applyMiddleware(topic.Handler(), middlewares...)
				r := h(msg)
				err = r.Error()
				return
			})
		case <-session.Context().Done():
			return nil
		}
	}
}

func (s *subscriber) handleError() {
	for err := range s.consumerGroup.Errors() {
		s.sendErrorCb(err)
	}
}

func (s *subscriber) sendErrorCb(err error) {
	if s.errCb != nil {
		s.errCb(err)
	}
}

func (s *subscriber) populateTopics(topics []messaging.Topic) []string {
	topicNames := make([]string, len(topics))
	for i, topic := range topics {
		if topic.IsDelay() {
			s.handlers[topic.Topic()] = newDelayedHandler(topic, s.consumerGroup)
		} else {
			s.handlers[topic.Topic()] = newNormalHandler(topic)
		}
		topicNames[i] = topic.Topic()
	}
	return topicNames
}

func (s *subscriberHandler) applyMiddleware(h messaging.SubscriptionHandler, middlewares ...messaging.MiddlewareFunc) messaging.SubscriptionHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
