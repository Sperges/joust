package db

import (
	"fmt"
	"joust/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func LoadWorldDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("failed to open splite: %w", err)
	}

	db.AutoMigrate(
		&model.BoxScore{},
		&model.HorseStats{},
		&model.Horse{},
		&model.KnightStats{},
		&model.Knight{},
		&model.Match{},
		&model.User{},
	)

	return db, nil
}
