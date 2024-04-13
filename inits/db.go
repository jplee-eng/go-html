package inits

import (
	"os"
	"sync"

	"github.com/jasonpatricklee/gowebserver/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

func init() {
	once.Do(func() {
		dsn := os.Getenv("DB_CONNECTION_STRING")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		utils.HandleError(err)
	})
}
func GetDb() *gorm.DB {
	return db
}
