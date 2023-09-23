package config

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	Logger        *logrus.Entry
	loggerRunOnce sync.Once
)

func NewLogger() {
	loggerRunOnce.Do(func() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		Logger = logrus.WithFields(
			logrus.Fields{
				"timestamp": time.Now(),
			},
		)
		logrus.SetOutput(os.Stdout)
		level, err := logrus.ParseLevel("info")
		if err != nil {
			level = logrus.InfoLevel
		}
		logrus.SetLevel(level)
	})

}
