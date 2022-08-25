package DataBase

import (
	"dogegambling/handlers"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgredb() *gorm.DB {
	dbURL := "postgres://postgres:docker@localhost:5432/gambelDB"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&handlers.User{}, &handlers.Payment{})
	return db
}
