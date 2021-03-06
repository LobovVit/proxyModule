package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"proxymodule/utils/config"
)

// TODO настроить логгер
func LogInit(Config *config.Cfg) (*zap.Logger, error) {

	//encoderCfg := zapcore.EncoderConfig{
	//	TimeKey:        "TIME",
	//	LevelKey:       "L",
	//	NameKey:        "N",
	//	CallerKey:      "C",
	//	FunctionKey:    zapcore.OmitKey,
	//	MessageKey:     "M",
	//	StacktraceKey:  "S",
	//	LineEnding:     zapcore.DefaultLineEnding,
	//	EncodeLevel:    zapcore.CapitalLevelEncoder,
	//	EncodeTime:     zapcore.ISO8601TimeEncoder,
	//	EncodeDuration: zapcore.StringDurationEncoder,
	//	EncodeCaller:   zapcore.ShortCallerEncoder,
	//}
	var level zapcore.Level
	switch Config.LOGLEVEL {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "INFO":
		level = zapcore.InfoLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "PANIC":
		level = zapcore.PanicLevel
	case "FATAL":
		level = zapcore.FatalLevel
	default:
		level = zapcore.DebugLevel
	}
	//core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, level)
	//logger := zap.New(core)
	//logger, err := zap.NewDevelopment() //NewDevelopment()
	cfg := zap.Config{
		Encoding:         "console", //"json",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	return logger, err
}
