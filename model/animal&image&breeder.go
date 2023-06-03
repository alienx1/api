package model

import "time"

type Animal struct {
	AnimalID   string    `json:"animal_id"`
	Birthday   time.Time `json:"birthday"`
	EarNumber  string    `json:"ear_number"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Gender     string    `json:"gender"`
	Color      string    `json:"color"`
	OwnerName  string    `json:"owner_name"`
	Defect     string    `json:"defect"`
	Image1     string    `json:"image1"`
	Image2     string    `json:"image2"`
	Image3     string    `json:"image3"`
	Image4     string    `json:"image4"`
	Image5     string    `json:"image5"`
	Breeder1   string    `json:"breeder1"`
	Breeder2   string    `json:"breeder2"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
	CraeteBy   string    `json:"create_by"`
	UpdateTime time.Time `json:"update_time"`
	UpdateBy   string    `json:"update_by"`
}
