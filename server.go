package main

import (
	"api/conf"
	"api/model/db"

	"api/router"
	"fmt"

	"github.com/gin-gonic/gin"
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
	r.Use(CORSMiddleware())
	r.Run(":80")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Continue handling the request
		c.Next()
	}
}
