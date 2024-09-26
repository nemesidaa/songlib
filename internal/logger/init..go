package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
}

func GetLogger() *logrus.Logger {
	return logger
}
