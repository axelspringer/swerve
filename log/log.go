package log

import (
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

// SetupLogger prepares the logger instance
func SetupLogger(level string, outType string) {
	logger.SetLevel(logrus.DebugLevel)
	if logLevel, err := logrus.ParseLevel(level); err == nil {
		logger.SetLevel(logLevel)
	}
	logger.WithField("a", "b")

	outType = strings.ToLower(outType)

	switch outType {
	case "json":
		logrus.SetFormatter(new(logrus.JSONFormatter))
	case "text":
		logrus.SetFormatter(new(logrus.TextFormatter))
	default:
		logrus.SetFormatter(new(logrus.TextFormatter))
	}
}

// Fatal wraps log.Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Fatalf wraps log.Fatalf
func Fatalf(fmt string, args ...interface{}) {
	logger.Fatalf(fmt, args...)
}

// Debug wraps log.Debug
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Debugf wraps log.Debugf
func Debugf(fmt string, args ...interface{}) {
	logger.Debugf(fmt, args...)
}

// Info wraps log.Info
func Info(args ...interface{}) {
	logger.Info(args)
}

// Infof wraps log.Infof
func Infof(fmt string, args ...interface{}) {
	logger.Infof(fmt, args...)
}

// Warn wraps log.Warn
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Warnf wraps log.Warnf
func Warnf(fmt string, args ...interface{}) {
	logger.Warnf(fmt, args...)
}

// Error wraps log.Error
func Error(args ...interface{}) {
	logger.Error(args)
}

// Errorf wraps log.Errorf
func Errorf(fmt string, args ...interface{}) {
	logger.Errorf(fmt, args...)
}

// ResponseWithMsg logs response specific data including error messages
func ResponseWithMsg(r *http.Request, code int, msg string) {
	logger.Infof("status=%d request=%s %s %s%s msg=%s", code, r.Proto, r.Method, r.Host, r.URL.Path, msg)
}

// Response logs response specific data
func Response(r *http.Request, code int) {
	logger.Infof("status=%d request=%s %s %s%s", code, r.Proto, r.Method, r.Host, r.URL.Path)
}
