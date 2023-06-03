package auth

import (
	"api/conf"
	"api/model"
	"api/model/db"
	"api/service"
	"api/service/key"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte
var hmacSecret []byte

func AdminLogin(c *gin.Context) {
	var res model.Login

	if err := c.ShouldBindJSON(&res); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Fail",
			"token":  nil,
		})
	} else {
		var a model.Admin
		conf.DB.Where("username = ?", res.Username).First(&a)
		if a.Username != res.Username || a.Status != 1 {
			c.JSON(http.StatusOK, gin.H{
				"status": "error",
				"error":  "User Dose Not Exists",
				"token":  nil,
			})
		}
		if a.Username == res.Username && a.Status == 1 {
			password := key.EncodeMd5(res.Password)
			if password == a.Password {
				hmacSampleSecret = []byte(os.Getenv("JWT_ANY_KEY"))
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": a.Username,
					"name":     a.Name,
					"rank":     a.Rank,
					"exp":      time.Now().Add(time.Hour * 3).Unix(),
				})
				tokenString, err := token.SignedString(hmacSampleSecret)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusOK, gin.H{
						"status": "error",
						"error":  err.Error(),
						"token":  nil,
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
						"error":  nil,
						"token":  tokenString,
					})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "error",
					"error":  "Login Failed",
					"token":  nil,
				})
			}
		}
	}
}

func UserLogin(c *gin.Context) {
	var res model.Login

	if err := c.ShouldBindJSON(&res); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Fail",
			"token":  nil,
		})
	} else {
		var u model.User
		conf.DB.Where("username = ?", res.Username).First(&u)
		if u.Username != res.Username || u.Status != 1 {
			c.JSON(http.StatusOK, gin.H{
				"status": "error",
				"error":  "User Dose Not Exists",
				"token":  nil,
			})
		}
		if u.Username == res.Username && u.Status == 1 {
			password := key.EncodeMd5(res.Password)
			if password == u.Password {
				hmacSampleSecret = []byte(os.Getenv("JWT_KEY"))
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": u.Username,
					"name":     u.Name,
					"exp":      time.Now().Add(time.Hour * 6).Unix(),
				})
				tokenString, err := token.SignedString(hmacSampleSecret)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusOK, gin.H{
						"status": "error",
						"error":  err.Error(),
						"token":  nil,
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
						"error":  nil,
						"token":  tokenString,
					})
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "error",
					"error":  "Login Failed",
					"token":  nil,
				})
			}
		}
	}
}

func Register(c *gin.Context) {
	var u db.User
	c.BindJSON(&u)
	fmt.Println(u.Address)
	u.Address = key.EncodeHex(u.Address)
	u.CreateTime = time.Now()
	u.Password = key.EncodeMd5(u.Password)
	u.ID = key.EncodeHex(u.ID)
	u.Credit = key.EncodeHex(u.Credit)
	if err := service.CreateUser(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"error":  nil,
		})
	}
}
