package controller

import (
	"api/model"
	"api/model/db"
	"api/service"
	"api/service/key"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindTableAdmin(c *gin.Context) {
	user := c.Param("user")
	name := c.Param("name")
	var a []db.Admin
	if user == "FULLTABLE" && name == "none" {
		if err := service.GetTableAdmin(&a); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, a)
		}
	}
	if user != "FULLTABLE" && name == "none" {
		if err := service.GetTableAdminByUser(&a, user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, a)
		}
	}
	if name != "none" && user == "none" {
		if err := service.GetTableAdminByName(&a, name); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, a)
		}
	}
}

func CreateAdmin(c *gin.Context) {
	var a db.Admin
	c.BindJSON(&a)
	a.CreateBy = c.MustGet("username").(string)
	a.CreateTime = time.Now()
	fmt.Println(a)
	a.Password = key.EncodeMd5(a.Password)

	if err := service.CraeteAdmin(&a); err != nil {
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
func UpdateAdmin(c *gin.Context) {
	var a model.Admin
	user := c.Param("user")
	c.BindJSON(&a)
	a.UpdateBy = c.MustGet("username").(string)
	a.UpdateTime = time.Now()
	fmt.Println(a.Password)
	a.Password = key.EncodeMd5(a.Password)
	fmt.Println(a.Password)
	if err := service.UpdateAdmin(&a, user); err != nil {
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
func DeleteAdmin(c *gin.Context) {
	var a db.Admin
	user := c.Param("user")
	if err := service.DeleteAdmin(&a, user); err != nil {
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
func BanAdmin(c *gin.Context) {
	var a model.Admin
	user := c.Param("user")
	status := c.Param("status")
	service.FindSingleAdmin(&a, user)
	a.UpdateBy = c.MustGet("username").(string)
	a.UpdateTime = time.Now()
	if status == "1" {
		a.Status = 0
		if err := service.UpdateAdmin(&a, user); err != nil {
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
	} else {
		a.Status = 1
		if err := service.UpdateAdmin(&a, user); err != nil {
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

}
