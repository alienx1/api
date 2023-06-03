package db

import "time"

type Admin struct {
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"not null"`
	Name       string `gorm:"primary_key"`
	Rank       string `gorm:"not null"`
	Status     int    `gorm:"not null"`
	CreateTime time.Time
	UpdateTime time.Time
	CreateBy   string
	UpdateBy   string
}
