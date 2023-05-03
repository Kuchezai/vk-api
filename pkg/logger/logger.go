package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func New(fileName string) (*Logger, error) {

	logFile, err := os.OpenFile("logs/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetOutput(logFile)

	return (&Logger{log: logger}), nil
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Warning(msg string) {
	l.log.Warning(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.log.Fatal(msg)
}
