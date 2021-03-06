package logutil

import (
	"log/syslog"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

const (
	// DevProfile for development profile
	DevProfile = "dev"
	// ProdProfile for production profile
	ProdProfile = "prod"
	// AppName the application name
	AppName = "user-service"
)

var (
	// Logger logrus custom instance
	Logger = logrus.New()
)

// InitLogger will init logger configuration
func InitLogger(profile string) {
	if profile == DevProfile {
		Logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})
	} else {
		Logger.SetFormatter(&logrus.JSONFormatter{})
	}

	// set syslog
	// this will connect to local syslog (Ex. "/dev/log" or "/var/run/syslog" or "/var/run/log")
	hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, AppName)

	if err == nil {
		Logger.AddHook(hook)
	} else {
		Logger.Errorf("error init local syslog %s", err.Error())
	}
}

// Debug will log all event in debug mode
func Debug(message, event, key string) {
	entry := setFields(event, key)
	entry.Debug(message)
}

// Info will log all event in info mode
func Info(message, event, key string) {
	entry := setFields(event, key)
	entry.Info(message)
}

// Warn will log all event in warn mode
func Warn(message, event, key string) {
	entry := setFields(event, key)
	entry.Warn(message)
}

// Error will log all event in error mode
func Error(message, event, key string) {
	entry := setFields(event, key)
	entry.Error(message)
}

func setFields(event, key string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{
		"topic": AppName,
		"event": event,
		"key":   key,
	})
}
