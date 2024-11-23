package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/super-saga/go-x/messaging"
	"github.com/super-saga/go-x/messaging/kafka"
	messagingNr "github.com/super-saga/go-x/messaging/newrelic"
)

func Newrelic(name string, newRelicApp *newrelic.Application) messaging.MiddlewareFunc {
	return func(next messaging.SubscriptionHandler) messaging.SubscriptionHandler {
		return func(message messaging.Message) messaging.Response {
			if newRelicApp == nil {
				return next(message)
			}

			dt := messagingNr.DistributedTracingAttribute{
				Header: make(http.Header),
			}
			claim := message.GetMessageClaim().(kafka.MessageClaim)

			name = fmt.Sprintf("%s/Topic/%s", name, claim.Topic)
			for _, hdr := range claim.Headers {
				dt.Header.Add(string(hdr.Key), string(hdr.Value))
			}

			ctx := messagingNr.BeforeSubscribe(
				message.Context(), newRelicApp, name,
				dt,
			)
			defer messagingNr.AfterSubscribe(ctx)

			if tx := newrelic.FromContext(ctx); tx != nil {
				tx.AddAttribute("msg.partition", claim.Partition)
				tx.AddAttribute("msg.key", string(claim.Key))
				tx.AddAttribute("msg.offset", claim.Offset)
				tx.AddAttribute("msg.timestamp", claim.Timestamp.Format(time.RFC3339))
				ctx = newrelic.NewContext(ctx, tx)
			}

			r := next(message.WithContext(ctx))

			if r.IsError() {
				if r.Report() {
					messagingNr.NoticeError(ctx, r.Error())
				} else {
					messagingNr.NoticeExpectedError(ctx, r.Error())
				}
			}

			return r
		}
	}
}
