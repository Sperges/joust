package model

import "gorm.io/gorm"

type KnightStats struct {
	gorm.Model
	KnightID uint

	Power       int `json:"power"`
	Avoidance   int `json:"avoidance"`
	Technique   int `json:"technique"`
	Accuracy    int `json:"accuracy"`
	Balance     int `json:"balance"`
	Riding      int `json:"riding"`
	Deflection  int `json:"defense"`
	Valor       int `json:"valor"`
	Chivalry    int `json:"chivalry"`
	Skulduggery int `json:"skulduggery"`
	Showmanship int `json:"showmanship"`
}
