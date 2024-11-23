package messaging

import (
	"time"
)

type SubscriptionHandler func(message Message) Response
type MiddlewareFunc func(next SubscriptionHandler) SubscriptionHandler

type Topic interface {
	Topic() string
	Handler() SubscriptionHandler
	Middlewares() []MiddlewareFunc
	Codec() Codec
	IsDelay() bool
	Delay() time.Duration
	Concurrent() int64
}

type topicStream struct {
	topic       string
	handler     SubscriptionHandler
	middlewares []MiddlewareFunc
	codec       Codec
	delay       time.Duration
	concurrent  int64
}

func (g topicStream) Topic() string {
	return g.topic
}

func (g topicStream) Handler() SubscriptionHandler {
	return g.handler
}

func (g topicStream) Middlewares() []MiddlewareFunc {
	return g.middlewares
}

func (g topicStream) Codec() Codec {
	return g.codec
}

func (g topicStream) IsDelay() bool {
	return g.delay > 0
}

func (g topicStream) Delay() time.Duration {
	return g.delay
}

func (g topicStream) Concurrent() int64 {
	return g.concurrent
}

// WithDelayedTopic will add delay to the consumer for this topic.
func WithDelayedTopic(topic string, codec Codec, delay time.Duration, concurrent int64, handler SubscriptionHandler, middlewares ...MiddlewareFunc) Topic {
	return &topicStream{
		topic:       topic,
		handler:     handler,
		codec:       codec,
		delay:       delay,
		concurrent:  concurrent,
		middlewares: middlewares,
	}
}

// WithTopic will add the consumer for this topic.
func WithTopic(topic string, codec Codec, handler SubscriptionHandler, middlewares ...MiddlewareFunc) Topic {
	return &topicStream{
		topic:       topic,
		handler:     handler,
		codec:       codec,
		middlewares: middlewares,
	}
}
