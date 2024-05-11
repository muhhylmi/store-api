package logger

import (
	"github.com/sirupsen/logrus"
)

func Newlogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:            false, // Nonaktifkan warna
		DisableColors:          true,  // Matikan warna
		DisableTimestamp:       true,  // Nonaktifkan tanggal
		DisableLevelTruncation: true,  // Nonaktifkan pemotongan level
		FullTimestamp:          true,  // Tampilkan timestamp lengkap (jika DisableTimestamp=false)
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	return &Logger{
		Logger: logger,
	}
}

func (log *Logger) LogWithContext(ctx string, scope string) *logrus.Entry {
	return log.Logger.WithFields(logrus.Fields{
		"scope":   scope,
		"context": ctx,
	})
}
