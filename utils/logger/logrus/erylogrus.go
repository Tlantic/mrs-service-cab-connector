package erylogrus

import (
	"strings"

	"github.com/andrepinto/erygo"
	"github.com/sirupsen/logrus"
)

// Logger ...
type Logger struct {
	*logrus.Entry
}

// NewLogger ...
func NewLogger(e *logrus.Entry) *Logger {
	return &Logger{Entry: e}
}

// LogErr ...
func (er *Logger) LogErr(err error, erygoErr *erygo.Err) {
	er.WithError(err).Errorf("returning %+v", erygoErr)
}

// LogResp ...
func (er *Logger) LogResp(msg string, erygoResp *erygo.Response) {

}

// Log ...
func Log(err *erygo.Err, optionalLogger ...logrus.FieldLogger) *logrus.Entry {
	var logger = func() logrus.FieldLogger {
		if len(optionalLogger) > 0 {
			return optionalLogger[0]
		}
		return logrus.StandardLogger()
	}()
	fields := make(map[string]interface{}, len(err.Labels))
	for k, v := range err.Labels {
		fields[k] = v
	}
	entry := logger.WithFields(fields).
		WithField("kind", err.Info.Kind).
		WithField("service", err.Info.Service).
		WithField("status-http", err.StatusHTTP)
	if entry.Message == "" {
		entry.Message = err.Message + func() string {
			if len(err.Details) > 0 {
				return ": " + strings.Join(err.Details, ";")
			}
			return ""
		}()
	}
	return entry
}
