package main

import (
	"api/conf"
	"api/model/db"

	"api/router"
	"fmt"
	"time"

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))
	r.Run(":80")
}
