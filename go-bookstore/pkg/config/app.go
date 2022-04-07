package config

import (
	"log"
	"path/filepath"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {

	dbpath := getDbPath()
	log.Print("Opening database connection at path: ", dbpath)
	connection, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = connection
}

func getDbPath() string {
	_, running_file, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(running_file)

	dbpath := filepath.Join(basepath, "..", "db", "bookstore.db")
	return dbpath
}

func GetDb() *gorm.DB {
	return db
}
