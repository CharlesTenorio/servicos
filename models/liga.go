package models

import "gorm.io/gorm"

type Liga struct {
	gorm.Model
	Id              int    `json:"id"`
	Name            string `json:"Name"`
	Cc              string `json:"cc"`
	Has_leaguetable string `json:"has_leaguetable"`
	Has_toplist     string `json:"has_toplist"`
}
