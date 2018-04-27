package util

import (
	"time"

	"github.com/sirupsen/logrus"
)

var apiLogger *logrus.Logger

// InitLogger gera as instâncias de logger para a API e para os workers.
func InitLogger() {

	apiLogger = logrus.New()
	apiLogger.SetNoLock()

	apiLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	}

}

// GetLogger retorna a instância do logger da API
func GetLogger() *logrus.Entry {
	if apiLogger == nil {
		InitLogger()
	}

	return apiLogger.WithField("type", "json")
}
