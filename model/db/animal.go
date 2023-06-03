package db

import "time"

type Animal struct {
	AnimalID   string `gorm:"primary_key"`
	Birthday   time.Time
	EarNumber  string
	Name       string `gorm:"not null"`
	Type       string `gorn:"not null"`
	Gender     string `gorm:"not null"`
	Color      string `gorm:"not null"`
	OwnerName  string `gorm:"not null"`
	Defect     string
	Status     int `gorm:"default:0"`
	CreateTime time.Time
	CreateBy   string
	UpdateTime time.Time
	UpdateBy   string
}
