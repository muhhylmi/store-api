package logger

import "github.com/sirupsen/logrus"

type Logger struct {
	Logger *logrus.Logger
}

type LoggerInterface interface {
	LogWithContext(ctx string, scope string) *logrus.Entry
}
