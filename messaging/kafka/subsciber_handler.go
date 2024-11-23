package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/super-saga/go-x/messaging"
)

type consumerHandlerExecutor func(msg messaging.Message, topic messaging.Topic) error

type consumerHandler interface {
	Handle(
		session sarama.ConsumerGroupSession,
		message *sarama.ConsumerMessage,
		executor consumerHandlerExecutor,
	)
	Close()
	Begin(session sarama.ConsumerGroupSession)
	End(session sarama.ConsumerGroupSession)
}

type normalHandler struct {
	topic messaging.Topic
}

func newNormalHandler(topic messaging.Topic) consumerHandler {
	return &normalHandler{
		topic: topic,
	}
}

func (h normalHandler) Handle(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage, executor consumerHandlerExecutor) {
	ctx, cancel := context.WithCancel(session.Context())
	executor(&Message{
		ConsumerMessage: message,
		codec:           h.topic.Codec(),
		ctx:             ctx,
		cancel:          cancel,
	}, h.topic)
	session.MarkMessage(message, "")
	cancel()
}

func (h normalHandler) Close()                            {}
func (h normalHandler) Begin(sarama.ConsumerGroupSession) {}
func (h normalHandler) End(sarama.ConsumerGroupSession)   {}

type delayedHandler struct {
	topic         messaging.Topic
	consumerGroup sarama.ConsumerGroup
	commiterChan  chan delayedHandlerCommiter
	partitions    map[int32]*delayedHandlerPartition
	parLock       sync.RWMutex
}

type delayedHandlerPartition struct {
	inProcess int64
	wg        sync.WaitGroup
}

type delayedHandlerCommiter struct {
	session sarama.ConsumerGroupSession
	message *sarama.ConsumerMessage
	err     error
}

func newDelayedHandler(topic messaging.Topic, consumerGroup sarama.ConsumerGroup) consumerHandler {
	h := &delayedHandler{
		topic:         topic,
		consumerGroup: consumerGroup,
		commiterChan:  make(chan delayedHandlerCommiter),
	}

	go h.runCommiter()

	return h
}

func (h *delayedHandler) Handle(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage, executor consumerHandlerExecutor) {
	h.parLock.Lock()
	par := h.partitions[message.Partition]
	h.parLock.Unlock()

	if par.inProcess == h.topic.Concurrent() {
		par.wg.Wait()
		par.inProcess = 0
	}
	par.inProcess++

	if !h.delayedFunc(session, message) {
		return
	}

	par.wg.Add(1)
	go func() {
		defer par.wg.Done()
		ctx, cancel := context.WithCancel(session.Context())
		err := executor(&Message{
			ConsumerMessage: message,
			codec:           h.topic.Codec(),
			ctx:             ctx,
			cancel:          cancel,
		}, h.topic)
		h.commiterChan <- delayedHandlerCommiter{
			message: message,
			err:     err,
			session: session,
		}
	}()
}

func (h *delayedHandler) Close() {
	h.waitPartitions()
	close(h.commiterChan)
}

func (h *delayedHandler) Begin(session sarama.ConsumerGroupSession) {
	partitions := session.Claims()[h.topic.Topic()]
	h.partitions = make(map[int32]*delayedHandlerPartition)
	for _, partition := range partitions {
		h.partitions[partition] = &delayedHandlerPartition{}
	}
}

func (h *delayedHandler) End(sarama.ConsumerGroupSession) {
	h.waitPartitions()
}

func (h *delayedHandler) delayedFunc(session sarama.ConsumerGroupSession, message *sarama.ConsumerMessage) (accept bool) {
	pauser, resumer, sleeper := func() {}, func() {}, func() bool { return true }
	now := time.Now()
	diff := now.Sub(message.Timestamp)
	tolerate := 30 * time.Second
	diffNs := diff.Nanoseconds() - h.topic.Delay().Nanoseconds()
	partitions := map[string][]int32{
		message.Topic: {message.Partition},
	}

	if diffNs < 0 {
		diffNs *= -1
		if diffNs > tolerate.Nanoseconds() {
			pauser = func() {
				h.consumerGroup.Pause(partitions)
			}
			resumer = func() {
				h.consumerGroup.Resume(partitions)
			}
		}
		sleeper = func() bool {
			pauser()
			select {
			case <-session.Context().Done():
				return false
			case <-time.After(time.Duration(diffNs) * time.Nanosecond):
				resumer()
			}
			return true
		}
	}

	return sleeper()
}

func (h *delayedHandler) runCommiter() {
	for commit := range h.commiterChan {
		// TO DO: algorithm to define which message is valid to be commited
		commit.session.MarkMessage(commit.message, "")
	}
}

func (h *delayedHandler) waitPartitions() {
	var wg sync.WaitGroup
	for _, par := range h.partitions {
		wg.Add(1)
		go func(par *delayedHandlerPartition) {
			par.wg.Wait()
			wg.Done()
		}(par)
	}
	wg.Wait()
}
