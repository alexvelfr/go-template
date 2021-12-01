package logger

import (
	"context"

	logstashclientmicro "github.com/alexvelfr/logstash-client-micro"
	"github.com/gin-gonic/gin"
)

var logClient logstashclientmicro.Client

// InitLogger ...
func InitLogger(servceName, uri string, useInsecureSSL bool) {
	logClient = logstashclientmicro.NewClient(servceName, uri, useInsecureSSL)
}

// LogError log it
func LogError(reqID, action, file, data string, err error) {
	if logClient == nil {
		return
	}
	logClient.LogError(context.Background(), logstashclientmicro.Message{
		XReqID: reqID,
		Data:   data,
		File:   file,
		Action: action,
		Error:  err,
	})
}

func RecoveryLog(c *gin.Context, err interface{}) {
	err2, ok := err.(error)
	if ok {
		LogError("", "RECOVERY", "", "", err2)
	}
}
