package database

import (
	"github.com/go-sql-driver/mysql"
	"logur.dev/logur"
)

// SetLogger configures the global database logger.
func SetLogger(logger logur.Logger) {
	logger = logur.WithFields(logger, map[string]interface{}{"component": "mysql"})

	_ = mysql.SetLogger(logur.NewErrorPrintLogger(logger))
}
