package model

import "time"

type Admin struct {
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Rank       string    `json:"rank"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	CreateBy   string    `json:"create_by"`
	UpdateBy   string    `json:"update_by"`
}
