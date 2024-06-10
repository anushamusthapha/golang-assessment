package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	file, err := os.OpenFile("golang-assessment", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.SetOutput(file)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.SetLevel(logrus.DebugLevel)
	Log = log

}
