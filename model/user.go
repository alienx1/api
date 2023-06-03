package model

import "time"

type User struct {
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Birthday    time.Time `json:"birthday"`
	Ethnic      string    `json:"ethnic"`
	Nationality string    `json:"nationality"`
	Religion    string    `json:"religion"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Duration    time.Time `json:"duration"`
	ID          string    `json:"id"`
	Credit      string    `json:"credit"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	CreateBy    string    `json:"create_by"`
	UpdateBy    string    `json:"update_by"`
}
