package model

import "gorm.io/gorm"

type HorseStats struct {
	gorm.Model
	HorseID uint

	Speed      float64 `json:"speed"`
	Power      float64 `json:"power"`
	Temperment float64 `json:"temperment"`
}
