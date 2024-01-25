package data

import (
	"context"

	"gorm.io/gorm"
)

type HorseRepo struct {
	DB *gorm.DB
}

func (r *HorseRepo) RandomHorseName(ctx context.Context) (string, error) {
	return RandomFromTable(ctx, r.DB, "horse_names")
}

func (r *HorseRepo) RandomHorseBreed(ctx context.Context) (string, error) {
	return RandomFromTable(ctx, r.DB, "horse_breeds")
}
