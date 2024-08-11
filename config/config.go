package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeMySql()

	if err != nil {
		return fmt.Errorf("error initializing mysql: %w", err)
	}

	return nil
}

func GetMysql() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// initialize Logger
	logger = NewLogger(prefix)
	return logger
}
