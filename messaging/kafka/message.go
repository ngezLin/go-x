package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/super-saga/go-x/messaging"
)

type Message struct {
	*sarama.ConsumerMessage
	codec  messaging.Codec
	ctx    context.Context
	cancel context.CancelFunc
}

func (m *Message) Bind(input any) error {
	return m.codec.Decode(m.Value, input)
}

func (m *Message) Context() context.Context {
	return m.ctx
}

func (m *Message) WithContext(ctx context.Context) messaging.Message {
	if ctx == nil {
		ctx = context.Background()
	}

	r2 := new(Message)
	*r2 = *m
	r2.ctx = ctx
	return r2
}

func (m *Message) ContextCancel() {
	m.cancel()
}

type MessageClaim *sarama.ConsumerMessage

func (m *Message) GetMessageClaim() any {
	var messageClaim MessageClaim = m.ConsumerMessage
	return messageClaim
}
