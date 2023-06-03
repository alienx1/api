package db

import "time"

type User struct {
	Username    string `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Birthday    time.Time
	Ethnic      string `gorm:"not null"`
	Nationality string `gorm:"not null"`
	Religion    string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Duration    time.Time
	ID          string `gorm:"not null"`
	Credit      string
	Status      int
	CreateTime  time.Time
	UpdateTime  time.Time
	CreateBy    string
	UpdateBy    string
}
