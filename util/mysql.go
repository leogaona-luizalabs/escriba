package util

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

// OpenMySQLConnection devolve uma conexão válida com um banco MySQL.
func OpenMySQLConnection(dsn string) (*sql.DB, error) {
	logger := GetLogger().WithFields(logrus.Fields{
		"operation_name": "openMySQLConnection",
		"start":          time.Now().Format(time.RFC3339),
	})

	// formata o dsn do banco e tenta pegar o dsn pelo viper
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"status":     "error",
			"error":      err.Error(),
			"error_type": "open",
			"end":        time.Now().Format(time.RFC3339),
		}).Error()
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"status":     "error",
			"error":      err.Error(),
			"error_type": "ping",
			"end":        time.Now().Format(time.RFC3339),
		}).Error()
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"status": "OK",
		"end":    time.Now().Format(time.RFC3339),
	}).Info()

	return db, nil
}
