package log

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	xsync "github.com/super-saga/x/sync"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

/*
LogLevel:
debug
info
warn
error
dpanic
panic
fatal
*/

/*
logOption:
file
fileconsole
console
*/

const (
	DefaultLogger  = "default"
	defaultMaxSize = 1024
	sep            = string(os.PathSeparator)
)

var (
	Loggers      xsync.SyncMap[string, *zap.Logger]
	SugarLoggers xsync.SyncMap[string, *zap.SugaredLogger]
)

type logToOption int32

const (
	logToOptUndefined logToOption = iota
	logToFileOpt
	logToFileAndStdOut
	logToStdOutOpt
)

func strToLogOption(str string) logToOption {
	switch strings.ToLower(str) {
	case "file":
		return logToFileOpt
	case "stdout":
		return logToStdOutOpt
	case "filestdout", "stdoutfile":
		return logToFileAndStdOut
	default:
		return logToOptUndefined
	}
}

type logEnv int32

const (
	logEnvUndefined logEnv = iota
	logEnvLocal
	logEnvDev
	logEnvProd
)

func strToLogEnv(str string) logEnv {
	switch strings.ToLower(str) {
	case "local":
		return logEnvLocal
	case "dev":
		return logEnvDev
	case "prod":
		return logEnvProd
	default:
		return logEnvUndefined
	}
}

type LogConfig func(*zLogCfg)

// WithLogToOption to overwrite the default log destination.
// valid value are: file, stdout, filestdout, and stdoutfile; stdoutfile and filestdout behave the same way,
// it will log both to std out and file. Default value is stdout.
//
// The input will treat as case insensitive.
func WithLogToOption(logTo string) LogConfig {
	return func(zlc *zLogCfg) {
		if lto := strToLogOption(logTo); lto != logToOptUndefined {
			zlc.lto = lto
		}
	}
}

// WithLogToOption to overwrite the default log destination.
// Valid value are: local, dev, and prod. default value is local
//
// Use local for local development, use prod for production environment.
func WithLogEnvOption(env string) LogConfig {
	return func(zlc *zLogCfg) {
		if le := strToLogEnv(env); le != logEnvUndefined {
			zlc.le = le
		}
	}
}

// WithCaller will print from where the logging function is called.
//
// Default value is false because getting the caller information can affect performance,
// although a little
func WithCaller(b bool) LogConfig {
	return func(zlc *zLogCfg) {
		zlc.withCaller = b
	}
}

// AddCallerSkip will add caller skip if you set WithCaller true.
//
// Default value is 1, which means will skip this wrapper library.
func AddCallerSkip(skip int) LogConfig {
	return func(zlc *zLogCfg) {
		zlc.callerStackSkip += skip
	}
}

func DebugLogLevel() LogConfig {
	return func(zlc *zLogCfg) {
		zlc.ll = zap.DebugLevel.String()
	}
}

func InfoLogLevel() LogConfig {
	return func(zlc *zLogCfg) {
		zlc.ll = zap.InfoLevel.String()
	}
}

func WarnLogLevel() LogConfig {
	return func(zlc *zLogCfg) {
		zlc.ll = zap.WarnLevel.String()
	}
}

func ErrorLogLevel() LogConfig {
	return func(zlc *zLogCfg) {
		zlc.ll = zap.ErrorLevel.String()
	}
}

func PanicLogLevel() LogConfig {
	return func(zlc *zLogCfg) {
		zlc.ll = zap.PanicLevel.String()
	}
}

type zLogCfg struct {
	lto             logToOption
	le              logEnv
	serviceName     string
	withCaller      bool
	callerStackSkip int
	ll              string
}

func (cfg zLogCfg) zapOpts() []zap.Option {
	opts := []zap.Option{zap.Fields(zap.String("service-name", cfg.serviceName))}
	if cfg.withCaller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(cfg.callerStackSkip))
	}
	if env := cfg.le; env == logEnvDev || env == logEnvLocal {
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel), zap.Development())
	}
	return opts
}

// Init will initialize a logger. the function must be called at program entry point;
// or where the program inialize its dependencies before serving any request or processing
// any message.
func Init(serviceName string, opts ...LogConfig) {
	cfg := zLogCfg{
		serviceName:     serviceName,
		lto:             logToStdOutOpt,
		le:              logEnvLocal,
		ll:              "info",
		callerStackSkip: 1,
	}
	for _, o := range opts {
		o(&cfg)
	}

	writesTo := []zapcore.WriteSyncer{}
	switch cfg.lto {
	case logToFileOpt, logToFileAndStdOut:
		hook := &lumberjack.Logger{
			Filename:   getFileName(cfg.serviceName), // Log name
			MaxSize:    defaultMaxSize,               // File content size, MB
			MaxBackups: 7,                            // Maximum number of old files retained
			MaxAge:     30,                           // Maximum number of days to keep old files
			Compress:   true,                         // Is the file compressed
		}
		writesTo = append(writesTo, zapcore.AddSync(hook))
		if cfg.lto == logToFileOpt {
			break
		}
		fallthrough
	case logToStdOutOpt:
		writesTo = append(writesTo, zapcore.AddSync(os.Stdout))
	}

	logLevel, err := zap.ParseAtomicLevel(cfg.ll)
	if err != nil {
		log.Printf("error during set log level: %s, will use default value", err)
		logLevel = zap.NewAtomicLevel()
	}

	zcore := zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoder(cfg.le)),
		zap.CombineWriteSyncers(writesTo...),
		logLevel,
	)
	l := zap.New(zcore, cfg.zapOpts()...)

	Loggers.Store(DefaultLogger, l)
	SugarLoggers.Store(DefaultLogger, l.Sugar())
}

// InitForTest will init the logger and sugarLogger for testing purpose
// it will not print anything at all.
func InitForTest() {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLogger := zap.New(observedZapCore)
	Loggers.Store("", observedLogger)
	SugarLoggers.Store("", observedLogger.Sugar())
}

func getEncoder(env logEnv) zapcore.EncoderConfig {
	var encoderConfig zapcore.EncoderConfig

	switch env {
	case logEnvProd:
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		// encoderConfig = zapcore.EncoderConfig{
		// 	NameKey:       "logger",
		// 	StacktraceKey: "stacktrace",
		// }
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		// encoderConfig = zapcore.EncoderConfig{
		// 	NameKey:       "logger",
		// 	StacktraceKey: "stacktrace",
		// }
	}
	encoderConfig.LevelKey = "severity"
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "caller"
	encoderConfig.FunctionKey = "func"
	encoderConfig.EncodeLevel = encodeLevel()
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	return encoderConfig
}

func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}

func getFileName(appName string) string {
	suffix := time.Now().Format("2006-01-02T15-04-05.000000000")
	fileName := "storages" + sep + "logs" + sep + "log_" + suffix + "_" + appName + ".log"

	return fileName
}

func Sync() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		Loggers.Range(
			func(k string, l *zap.Logger) bool {
				l.Sync()
				return true
			},
		)
	}()
	go func() {
		defer wg.Done()
		SugarLoggers.Range(
			func(k string, s *zap.SugaredLogger) bool {
				s.Sync()
				return true
			},
		)
	}()
	wg.Wait()
}
