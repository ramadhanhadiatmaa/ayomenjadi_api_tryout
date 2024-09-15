package models

import "gorm.io/gorm"

type Tryout struct {
	gorm.Model
	Title  string `json:"title"`
	Api string `json:"api"`
}
