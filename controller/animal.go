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

func CreateAnimal(c *gin.Context) {
	var req model.Animal
	var a db.Animal
	var ai db.AnimalImage
	var b db.Breeder
	c.BindJSON(&req)
	fmt.Println(req.CraeteBy)
	// Add data to the Animal struct
	a.AnimalID = key.GenerateKeyAnimal(req.Type, req.Gender, req.Color, req.Birthday)
	a.Birthday = req.Birthday
	a.EarNumber = req.EarNumber
	a.Name = req.Name
	a.Type = req.Type
	a.Gender = req.Gender
	a.Color = req.Color
	a.OwnerName = req.OwnerName
	a.Defect = req.Defect
	a.CreateTime = time.Now()
	a.CreateBy = req.CraeteBy

	if a.CreateBy == "" {
		a.CreateBy = c.MustGet("username").(string)
	}

	if a.OwnerName == "" {
		a.OwnerName = c.MustGet("username").(string)
	}
	// Add data to the AnimalImage struct
	ai.AnimalID = a.AnimalID
	ai.Image1 = req.Image1
	ai.Image2 = req.Image2
	ai.Image3 = req.Image3
	ai.Image4 = req.Image4
	ai.Image5 = req.Image5
	ai.CreateTime = time.Now()
	ai.CreateBy = req.CraeteBy

	if ai.CreateBy == "" {
		ai.CreateBy = c.MustGet("username").(string)
	}
	// Add data to the Breeder struct
	b.AnimalID = a.AnimalID
	b.Breeder1 = req.Breeder1
	b.Breeder2 = req.Breeder2
	b.CreateBy = req.CraeteBy
	b.CreateTime = time.Now()
	if b.CreateBy == "" {
		b.CreateBy = c.MustGet("username").(string)
	}

	if err1 := service.CreateAnimal(&a); err1 != nil {
		fmt.Println(err1.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err1.Error(),
		})
	} else {
		if err2 := service.CreateBreeder(&b); err2 != nil {
			fmt.Println(err2.Error())

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err2.Error(),
			})
		} else {

			if err3 := service.CreateAnimalImg(&ai); err3 != nil {
				fmt.Println(err3.Error())

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status": "error",
					"error":  err3.Error(),
				})
				fmt.Println(err1.Error())
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"error":  nil,
				})
			}
		}
	}
}

func FindTableAnimal(c *gin.Context) {
	var res []model.Animal
	animal_id := c.Param("animal_id")
	owner_name := c.Param("owner_name")
	name := c.Param("name")
	Type := c.Param("type")
	if animal_id == "FULLTABLE" && owner_name == "none" && name == "none" && Type == "none" {
		if err := service.GetTableAnimal(&res); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {

			c.JSON(http.StatusOK, res)
		}
	} else if animal_id != "none" && animal_id != "FULLTABLE" && owner_name == "none" && name == "none" && Type == "none" {
		if err := service.GetTableAnimalByAnimalID(&res, animal_id); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, res)
		}
	} else if animal_id == "none" && owner_name != "none" && name == "none" && Type == "none" {
		if err := service.GetTableAnimalOwnerName(&res, owner_name); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, res)
		}
	} else if animal_id == "none" && owner_name == "none" && name != "none" && Type == "none" {
		if err := service.GetTableAnimalName(&res, name); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, res)
		}
	} else if animal_id == "none" && owner_name == "none" && name == "none" && Type != "none" {
		if err := service.GetTableAnimalType(&res, Type); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, res)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Not Fond",
		})
	}

}

func GetTableAnimalByAnimalIDWithPet(c *gin.Context) {
	var res model.Animal
	animal_id := c.Param("animal_id")
	if err := service.GetTableAnimalByAnimalIDWithPet(&res, animal_id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
