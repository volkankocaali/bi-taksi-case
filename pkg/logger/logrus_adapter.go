package logger

import "github.com/sirupsen/logrus"

type LogrusAdapter struct {
	logger *logrus.Logger
}

func NewLogrusAdapter() *LogrusAdapter {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &LogrusAdapter{logger: logger}
}

func (l *LogrusAdapter) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusAdapter) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogrusAdapter) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LogrusAdapter) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *LogrusAdapter) Println(v ...interface{}) {
	l.logger.Println(v...)
}
