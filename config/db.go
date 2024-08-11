package config

import (
	"github.com/RenatoFa/gopportunities.git/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySql() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	// Create DB and Connect
	dsn := "root@tcp(127.0.0.1:3306)/gopportunities?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
