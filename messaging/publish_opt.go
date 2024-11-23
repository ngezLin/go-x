package messaging

import "github.com/super-saga/go-x/messaging/codec"

type PublishOption func(*publishOption)

type publishOption struct {
	codec Codec
}

func NewPublishOption() *publishOption {
	return &publishOption{}
}

func (p *publishOption) ApplyOption(opts ...PublishOption) {
	for _, opt := range opts {
		opt(p)
	}
}

func (p publishOption) Codec() Codec {
	return p.codec
}

func WithCodec(codec Codec) PublishOption {
	return func(po *publishOption) {
		po.codec = codec
	}
}

func WithDefaultCodec() PublishOption {
	return WithCodec(codec.NewJson("1.0"))
}
