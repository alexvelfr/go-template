package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	os.Mkdir(viper.GetString("app.log.dir"), os.ModeAppend)
	file, err := os.OpenFile(viper.GetString("app.log.file"), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(file)
}

// LogError log it
func LogError(action, file, data string, err error) {
	logrus.WithFields(
		logrus.Fields{
			"action": action,
			"file":   file,
			"data":   data,
		},
	).Error(err)
}
