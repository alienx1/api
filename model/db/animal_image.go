package db

import "time"

type AnimalImage struct {
	AnimalID   string
	Image1     string `gorm:"type:longtext"`
	Image2     string `gorm:"type:longtext"`
	Image3     string `gorm:"type:longtext"`
	Image4     string `gorm:"type:longtext"`
	Image5     string `gorm:"type:longtext"`
	CreateTime time.Time
	CreateBy   string
	UpdateTime time.Time
	UpdateBy   string
}
