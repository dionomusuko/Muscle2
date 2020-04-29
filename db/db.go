package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func NewDB() *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		"user",
		"password",
		"localhost",
		"3306",
		"go-db",
	)

	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
	err = conn.DB().Ping()

	if nil != err {
		panic(err)
	}
	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
