package models

import "gorm.io/gorm"

type Time struct {
	gorm.Model
	Id    int64  `json:"id"`
	Name  string `json:"Name"`
	IdImg int64  `json:"idImg"`
	Cc    string `json:"cc"`
}
