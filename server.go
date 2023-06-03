package main

import (
	"api/conf"
	"api/model/db"

	"api/router"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	conf.DB, err = gorm.Open("mysql", conf.DbURL(conf.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer conf.DB.Close()
	conf.DB.AutoMigrate(
		&db.Admin{},
		&db.User{},
		&db.Animal{},
		&db.AnimalImage{},
		&db.Breeder{},
	)
	r := router.Router()
	r.Use(cors.Default())
	r.Run(":80")
}
