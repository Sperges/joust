package model

import "gorm.io/gorm"

type Knight struct {
	gorm.Model
	KnightStats KnightStats
	UserID      uint `json:"user_id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint   `json:"age"`
	Region    string `json:"region"`
}
