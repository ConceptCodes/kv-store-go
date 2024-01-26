package sqlite

import (
	"sync"

	"gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDBInstance() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	})
	return db, err
}
