package log

import (
	logrus "github.com/Sirupsen/logrus"
)

func log(args map[string]string) *logrus.Entry {
	fields := logrus.Fields{}
	for k, v := range args {
		fields[k] = v
	}
	return logrus.WithFields(fields)
}

func Debug(msg string, args map[string]string) {
	log(args).Debug(msg)
}

func Info(msg string, args map[string]string) {
	log(args).Debug(msg)
}

func Warn(msg string, args map[string]string) {
	log(args).Warn(msg)
}

func Error(msg string, args map[string]string) {
	log(args).Panic(msg)
}
