package model

import "gorm.io/gorm"

type Match struct {
	gorm.Model

	BoxScores []BoxScore

	HomeId    uint `json:"home_id"`
	HomeTotal uint `json:"home_total"`

	AwayId    uint `json:"away_id"`
	AwayTotal uint `json:"away_total"`

	WinnerId uint `json:"winner_id"`
}
