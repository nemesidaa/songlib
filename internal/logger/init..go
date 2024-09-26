package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger(outputdest, level string) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	dst, err := parseDest(outputdest)
	if err != nil {
		panic(err) // ну не хочу я ни проводить ошибки до main-а, ни скипать ошибку.
	}
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
		ForceColors:   true,
	})
	logger.SetLevel(lvl)
	logger.SetOutput(dst)
}

func GetLogger() *logrus.Logger {
	return logger
}

// internalCtx

func parseDest(outputdest string) (io.Writer, error) {
	switch outputdest {
	case "stdout":
		return os.Stdout, nil
	default:
		return os.OpenFile(outputdest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
}
