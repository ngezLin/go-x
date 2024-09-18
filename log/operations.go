package log

import (
	"context"
	"fmt"

	"github.com/super-saga/x/log/ctxdata"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func populateContextFields(ctx context.Context) []Field {
	return []Field{
		String(LogKeyCorrelationId, ctxdata.GetCorrelationId(ctx)),

		String(LogKeyTraceparent, ctxdata.GetTraceParent(ctx)),
		String(LogKeyTraceID, ctxdata.GetTraceID(ctx)),
		String(LogKeySpanID, ctxdata.GetSpanID(ctx)),
		Bool(LogKeyTraceSampled, ctxdata.GetTraceSampled(ctx)),

		String(LogKeyUserAgent, ctxdata.GetUserAgent(ctx)),
		String(LogKeyHost, ctxdata.GetHost(ctx)),
		String(LogKeyIP, ctxdata.GetIP(ctx)),
		String(LogKeyForwardedFor, ctxdata.GetForwardedFor(ctx)),
		String(LogKeyPid, ctxdata.GetPid(ctx)),
	}
}

type Field = zapcore.Field
type ObjectEncoder = zapcore.ObjectEncoder

var (
	Any        = zap.Any
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	Int        = zap.Int
	Intp       = zap.Intp
	Ints       = zap.Ints
	Int8       = zap.Int8
	Int8s      = zap.Int8s
	Int8p      = zap.Int8p
	Int16      = zap.Int16
	Int16p     = zap.Int16p
	Int16s     = zap.Int16s
	Int32      = zap.Int32
	Int32p     = zap.Int32p
	Int32s     = zap.Int32s
	Int64      = zap.Int64
	Int64p     = zap.Int16p
	Int64s     = zap.Int16s
	Float32    = zap.Float32
	Float32s   = zap.Float32s
	Float32p   = zap.Float32p
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float64s   = zap.Float64s
	Uint       = zap.Uint
	Uints      = zap.Uints
	Uintp      = zap.Uintp
	Uint8      = zap.Uint8
	Uint8s     = zap.Uint8s
	Uint8p     = zap.Uint8p
	Uint16     = zap.Uint16
	Uint16s    = zap.Uint16s
	Uint16p    = zap.Uint16p
	Uint32     = zap.Uint32
	Uint32s    = zap.Uint32s
	Uint32p    = zap.Uint32p
	Uint64     = zap.Uint64
	Uint64s    = zap.Uint64s
	Uint64p    = zap.Uint64p
	String     = zap.String
	Strings    = zap.Strings
	Stringp    = zap.Stringp
	Stringer   = zap.Stringer
	Skip       = zap.Skip
	Time       = zap.Time
	Timep      = zap.Timep
	Times      = zap.Times
	Duration   = zap.Duration
	Durationp  = zap.Durationp
	Durations  = zap.Durations
	Err        = zap.Error
	Errs       = zap.Errors
	Namespace  = zap.Namespace
	Object     = zap.Object
)

func WithFields(ctx context.Context, fields ...Field) *zap.Logger {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return nil
	}

	return logger.With(populateContextFields(ctx)...).With(fields...)
}

func Info(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Info(message, fields...)
}

func Infof(ctx context.Context, message string, i ...interface{}) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Info(fmt.Sprintf(message, i...))
}

func Debug(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Debug(message, fields...)
}

func Debugf(ctx context.Context, message string, i ...interface{}) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Debug(fmt.Sprintf(message, i...))
}

func Warn(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Warn(message, fields...)
}

func Warnf(ctx context.Context, message string, i ...interface{}) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Warn(fmt.Sprintf(message, i...))
}

func Error(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Error(message, fields...)
}

func Errorf(ctx context.Context, message string, i ...interface{}) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Error(fmt.Sprintf(message, i...))
}

func Fatal(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Fatal(message, fields...)
}

func Fatalf(ctx context.Context, message string, i ...interface{}) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Fatal(fmt.Sprintf(message, i...))
}

func Panic(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).Panic(message, fields...)
}

func DPanic(ctx context.Context, message string, fields ...Field) {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return
	}
	logger.With(populateContextFields(ctx)...).DPanic(message, fields...)
}

func With(ctx context.Context, fields ...Field) *zap.Logger {
	logger, ok := Loggers.Load(DefaultLogger)
	if !ok {
		return nil
	}
	return logger.With(populateContextFields(ctx)...).With(fields...)
}
