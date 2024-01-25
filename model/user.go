package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Password string `json:"password"`

	ActiveKnightID uint
	Knights        []Knight
	ActiveHorseID  uint `json:"active_horse_id"`
	Horses         []Horse
}
