package middleware

import (
	"time"

	xlog "github.com/super-saga/go-x/log"
	"github.com/super-saga/go-x/messaging"
	"github.com/super-saga/go-x/messaging/kafka"
)

func Log(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
	return func(message messaging.Message) messaging.Response {
		startTime := time.Now()
		claim := message.GetMessageClaim().(kafka.MessageClaim)

		xlog.Info(message.Context(), "start consuming kafka message",
			xlog.String("kafka-topic", claim.Topic),
			xlog.Int64("kafka-offset", claim.Offset),
			xlog.Int32("kafka-partition", claim.Partition),
			xlog.String("kafka-key", string(claim.Key)),
			xlog.String("kafka-timestamp", claim.Timestamp.Format(time.RFC3339)),
		)

		r := next(message)

		xlog.Info(message.Context(), "finish consuming kafka message",
			xlog.String("start-time", startTime.Format(time.RFC3339)),
			xlog.Any("response", r),
			xlog.Int64("latency-ms", time.Since(startTime).Milliseconds()),
		)

		return r
	}
}
