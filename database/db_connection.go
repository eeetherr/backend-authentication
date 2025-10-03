// package database

// import (
// 	"fmt"
// 	"log"
// 	"sync"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var (
// 	db   *gorm.DB
// 	once sync.Once
// )

// func GetDB() *gorm.DB {
// 	once.Do(func() {
// 		dsn := "host=localhost user=postgres password=psg dbname=postgres port=5432 sslmode=disable"
// 		var err error
// 		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 			log.Fatal("Failed to connect to database:", err)
// 		}

// 		fmt.Println("Database connected successfully (GORM)")
// 	})
// 	return db
// }

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB = &gorm.DB{}

func InitDB() {
	dsn := "host=localhost user=jayant.m password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("✅ Connected to database")
}

func GetDB() *gorm.DB {
	return DB
}
