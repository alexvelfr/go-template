package logger

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var out *os.File

// InitLogger ...
func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	os.Mkdir(viper.GetString("app.log.dir"), 0755)
	filepath := path.Join(viper.GetString("app.log.dir"), viper.GetString("app.log.file"))
	out, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(out)
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

func GetOutFile() *os.File {
	return out
}
