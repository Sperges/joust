package model

import "gorm.io/gorm"

type Horse struct {
	gorm.Model
	HorseStats HorseStats
	UserID     uint

	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
}
