package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapAdapter struct {
	logger *zap.Logger
}

func NewZapAdapter() *ZapAdapter {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	return &ZapAdapter{
		logger: logger,
	}
}

func (z *ZapAdapter) Info(args ...interface{}) {
	z.logger.Info("", zap.Any("info", args))
}

func (z *ZapAdapter) Error(args ...interface{}) {
	z.logger.Error("", zap.Any("error", args))
}

func (z *ZapAdapter) Fatal(args ...interface{}) {
	z.logger.Fatal("", zap.Any("fatal", args))
}

func (z *ZapAdapter) Printf(format string, v ...interface{}) {
	z.logger.Sugar().Infof(format, v...)
}

func (z *ZapAdapter) Println(v ...interface{}) {
	z.logger.Sugar().Info(v...)
}
