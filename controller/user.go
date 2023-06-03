package controller

import (
	"api/model"
	"api/model/db"
	"api/service"
	"api/service/key"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func FindTableUser(c *gin.Context) {
	var u []model.User
	user := c.Param("user")
	name := c.Param("name")
	phone := c.Param("phone")
	id_card := c.Param("id_card")
	if user == "FULLTABLE" {
		if err := service.GetTableUser(&u); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			for i := 0; i < len(u); i++ {
				u[i].ID = key.DecodeHex(u[i].ID)
				u[i].Address = key.DecodeHex(u[i].Address)
				u[i].Credit = key.DecodeHex(u[i].Credit)
			}
			c.JSON(http.StatusOK, u)

		}
	}
	if user != "FULLTABLE" && user != "none" {
		if err := service.GetTableUserByUser(&u, user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {

			for i := 0; i < len(u); i++ {
				u[i].ID = key.DecodeHex(u[i].ID)
				u[i].Address = key.DecodeHex(u[i].Address)
				u[i].Credit = key.DecodeHex(u[i].Credit)
			}
			c.JSON(http.StatusOK, u)
		}
	}
	if name != "none" && user != "FULLTABLE" {
		if err := service.GetTableUserByName(&u, name); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			for i := 0; i < len(u); i++ {
				u[i].ID = key.DecodeHex(u[i].ID)
				u[i].Address = key.DecodeHex(u[i].Address)
				u[i].Credit = key.DecodeHex(u[i].Credit)
			}
			c.JSON(http.StatusOK, u)
		}
	}
	if phone != "none" && user != "FULLTABLE" {
		if err := service.GetTableUserByPhone(&u, phone); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			for i := 0; i < len(u); i++ {
				u[i].ID = key.DecodeHex(u[i].ID)
				u[i].Credit = key.DecodeHex(u[i].Credit)
			}
			c.JSON(http.StatusOK, u)
		}
	}
	if id_card != "none" && user != "FULLTABLE" {
		if err := service.GetTableUserByName(&u, id_card); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			for i := 0; i < len(u); i++ {
				u[i].ID = key.DecodeHex(u[i].ID)
				u[i].Address = key.DecodeHex(u[i].Address)
				u[i].Credit = key.DecodeHex(u[i].Credit)
			}
			c.JSON(http.StatusOK, u)
		}
	}
}

func CreateUser(c *gin.Context) {
	var u db.User
	c.BindJSON(&u)
	fmt.Println(u.Address)
	u.Status = 1
	u.CreateBy = c.MustGet("username").(string)
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

func UpdateUser(c *gin.Context) {
	var u model.User
	user := c.Param("user")
	id := c.Param("id")
	c.BindJSON(&u)
	u.UpdateBy = c.MustGet("username").(string)
	u.UpdateTime = time.Now()
	u.Password = key.EncodeMd5(u.Password)
	u.ID = key.EncodeHex(u.ID)
	u.Address = key.EncodeHex(u.Address)
	u.Credit = key.EncodeHex(u.Credit)
	if err := service.UpdateUser(&u, user, id); err != nil {
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

func DeleteUser(c *gin.Context) {
	var u model.User
	user := c.Param("user")

	if err := service.DeleteUser(&u, user); err != nil {
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

func BanUser(c *gin.Context) {
	var u model.User
	user := c.Param("user")
	service.FindSingleUser(&u, user)
	u.UpdateBy = c.MustGet("username").(string)
	u.UpdateTime = time.Now()
	if u.Status == 1 {
		u.Status = 0
	} else {
		u.Status = 1
	}
	fmt.Println(u.Status)
	if err := service.UpdateStatus(&u, user); err != nil {
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

func AddAirTime(c *gin.Context) {
	var u model.User
	user := c.Param("user")
	duration := c.Param("airtime")
	fmt.Println(duration)
	var d time.Duration
	var err error
	if strings.Contains(duration, "d") {
		parts := strings.Split(duration, "d")
		if len(parts) == 2 {
			days, _ := strconv.Atoi(parts[0])
			hours, _ := strconv.Atoi(parts[1])
			d = time.Duration(days*24+hours) * time.Hour
		} else {
			err = errors.New("Invalid duration format")
		}
	} else {
		d, err = time.ParseDuration(duration)
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		u.Duration = time.Now().Add(d)
		u.UpdateBy = c.MustGet("username").(string)
		u.UpdateTime = time.Now()

		if err1 := service.UpdateDuration(&u, user); err1 != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err1.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"error":  nil,
			})
		}
	}
}
func UserProfile(c *gin.Context) {
	var u model.User
	user := c.Param("user")
	if user == "none" {
		user = c.MustGet("username").(string)
	}
	fmt.Println(user)
	if err := service.FindSingleUser(&u, user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "ok",
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, u)
	}
}
