package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

type TixLogger struct {
	Hostname string
	*logrus.Logger
}

func MustGetLogger() *TixLogger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	tixLogger := &TixLogger{hostname, logrus.StandardLogger()}
	return tixLogger
}

func (tx *TixLogger) Info(args ...interface{}) {
	fields(tx).Info(args...)
}

func (tx *TixLogger) Infof(format string, args ...interface{}) {
	fields(tx).Infof(format, args...)
}

func (tx *TixLogger) Debug(args ...interface{}) {
	fields(tx).Debug(args...)
}

func (tx *TixLogger) Debugf(format string, args ...interface{}) {
	fields(tx).Debugf(format, args...)
}

func (tx *TixLogger) Warn(args ...interface{}) {
	fields(tx).Warn(args...)
}

func (tx *TixLogger) Warnf(format string, args ...interface{}) {
	fields(tx).Warnf(format, args...)
}

func (tx *TixLogger) Error(args ...interface{}) {
	fields(tx).Error(args...)
}

func (tx *TixLogger) Errorf(format string, args ...interface{}) {
	fields(tx).Errorf(format, args...)
}

// DebugfWithId write formatted debug level log with added log_id field
func (tx *TixLogger) DebugfWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Debugf(format, args...)
}

// InfofWithId write formatted info level log with added log_id field
func (tx *TixLogger) InfofWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Infof(format, args...)
}

// InfoWithId write info level log with added log_id field
func (tx *TixLogger) InfoWithId(id string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Info(args...)
}

// ErrorfWithId write formatted error level log with added log_id field
func (tx *TixLogger) ErrorfWithId(id string, format string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Errorf(format, args...)
}

// ErrorWithId write error level log with added log_id field
func (tx *TixLogger) ErrorWithId(id string, args ...interface{}) {
	fields(tx).WithField("log_id", id).Error(args...)
}

func fields(tx *TixLogger) *logrus.Entry {
	file, line := getCaller()
	return tx.Logger.WithField("time", time.Now().UTC().Format(time.RFC3339)).WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return file, line
}
