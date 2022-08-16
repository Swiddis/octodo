package model

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(database string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch database {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open("postgresql://todo:password@postgres:5432/todo"), &gorm.Config{})
	default:
		return nil, errors.New("unimplemented database")
	}
	db.AutoMigrate(&Todo{})
	return db, err
}
