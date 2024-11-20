package inits

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

type DBConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func InitDB() error {
	var initError error

	once.Do(func() {
		dsn := os.Getenv("DB_CONNECTION_STRING")
		if dsn == "" {
			initError = fmt.Errorf("database connection string not found in environment variables")
			return
		}

		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}

		var dbConn *gorm.DB
		dbConn, initError = gorm.Open(postgres.Open(dsn), gormConfig)
		if initError != nil {
			return
		}

		sqlDB, err := dbConn.DB()
		if err != nil {
			initError = fmt.Errorf("failed to get underlying *sql.DB: %w", err)
			return
		}

		config := DBConfig{
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: time.Hour,
		}

		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

		if err := sqlDB.Ping(); err != nil {
			initError = fmt.Errorf("failed to ping database: %w", err)
			return
		}

		db = dbConn
	})

	return initError
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}
	return db, nil
}

func CloseDB() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying *sql.DB: %w", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}
