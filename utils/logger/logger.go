package logger

import (
	"go.uber.org/zap"
)

// TODO настроить логгер
func LogInit() (*zap.Logger, error) {

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
	//core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zapcore.DebugLevel)
	//logger := zap.New(core)

	logger, err := zap.NewDevelopment() //NewDevelopment()
	return logger, err
}
