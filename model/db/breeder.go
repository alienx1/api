package db

import "time"

type Breeder struct {
	AnimalID   string
	Breeder1   string
	Breeder2   string
	CreateTime time.Time
	CreateBy   string
	UpdateTime time.Time
	UpdateBy   string
}
