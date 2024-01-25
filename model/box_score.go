package model

import "gorm.io/gorm"

type BoxScore struct {
	gorm.Model

	MatchID  uint
	KnightID uint

	PassNumber int
	Score      int

	AccuracyRoll   int
	AvoidanceRoll  int
	PowerRoll      int
	DeflectionRoll int
	TechniqueRoll  int
	BalanceRoll    int
}
