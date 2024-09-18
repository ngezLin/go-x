package audit

import (
	"context"
	"fmt"
	"strings"

	"github.com/super-saga/x/log"
	"github.com/super-saga/x/log/ctxdata"
)

func populateContextFields(ctx context.Context) []log.Field {
	return []log.Field{
		log.String(log.LogKeyTraceparent, ctxdata.GetTraceParent(ctx)),
		log.String(log.LogKeyTraceID, ctxdata.GetTraceID(ctx)),
		log.String(log.LogKeySpanID, ctxdata.GetSpanID(ctx)),
		log.Bool(log.LogKeyTraceSampled, ctxdata.GetTraceSampled(ctx)),

		log.Namespace(audit),
		log.String(log.LogKeyUserAgent, ctxdata.GetUserAgent(ctx)),
		log.String(log.LogKeyHost, ctxdata.GetHost(ctx)),
		log.String(log.LogKeyIP, ctxdata.GetIP(ctx)),
		log.String(log.LogKeyForwardedFor, ctxdata.GetForwardedFor(ctx)),
		log.String(log.LogKeyPid, ctxdata.GetPid(ctx)),
	}
}

const (
	audit               = "audit"
	logKeyClientAppName = "client-app-name"
	logKeyUserId        = "user-id"
	logKeyActivityData  = "activity-data"
)

var logMessage = fmt.Sprintf("[%s]", strings.ToUpper(audit))

type Message struct {
	ClientAppName string
	UserId        string
	ActivityData  interface{}
}

func Info(ctx context.Context, message Message) {
	logger, ok := log.Loggers.Load(log.DefaultLogger)
	if !ok {
		return
	}
	logger.Named(audit).With(populateContextFields(ctx)...).With([]log.Field{
		log.String(logKeyClientAppName, message.ClientAppName),
		log.String(logKeyUserId, message.UserId),
	}...).Info(logMessage, log.Any(logKeyActivityData, message.ActivityData))
}

func Debug(ctx context.Context, message Message) {
	logger, ok := log.Loggers.Load(log.DefaultLogger)
	if !ok {
		return
	}
	logger.Named(audit).With(populateContextFields(ctx)...).With([]log.Field{
		log.String(logKeyClientAppName, message.ClientAppName),
		log.String(logKeyUserId, message.UserId),
	}...).Debug(logMessage, log.Any(logKeyActivityData, message.ActivityData))
}

func Warn(ctx context.Context, message Message) {
	logger, ok := log.Loggers.Load(log.DefaultLogger)
	if !ok {
		return
	}
	logger.Named(audit).With(populateContextFields(ctx)...).With([]log.Field{
		log.String(logKeyClientAppName, message.ClientAppName),
		log.String(logKeyUserId, message.UserId),
	}...).Warn(logMessage, log.Any(logKeyActivityData, message.ActivityData))
}

func Error(ctx context.Context, message Message) {
	logger, ok := log.Loggers.Load(log.DefaultLogger)
	if !ok {
		return
	}
	logger.Named(audit).With(populateContextFields(ctx)...).With([]log.Field{
		log.String(logKeyClientAppName, message.ClientAppName),
		log.String(logKeyUserId, message.UserId),
	}...).Error(logMessage, log.Any(logKeyActivityData, message.ActivityData))
}
