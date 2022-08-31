package DataBase

import (
	"dogegambling/handlers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgredb(user, pass, server, port, dbname string) *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, pass, server, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&handlers.User{}, &handlers.Payment{}, &handlers.Bet{})
	return db
}
