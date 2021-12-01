package logger

import (
	"context"
	"errors"

	logstashclientmicro "github.com/alexvelfr/logstash-client-micro"
)

type LogWriter struct{}

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

func (l LogWriter) Write(p []byte) (int, error) {
	LogError("", "PANIC", "", string(p), errors.New("PANIC"))
	return len(p), nil
}
