package models

import "gorm.io/gorm"

type Tryout struct {
	gorm.Model
	To  string `json:"to"`
	Api string `json:"api"`
}
