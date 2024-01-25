package data

import (
	"context"
	"joust/model"

	"gorm.io/gorm"
)

type MatchRepo struct {
	DB *gorm.DB
}

func (r *MatchRepo) Create(ctx context.Context, match *model.Match) error {
	if err := r.DB.WithContext(ctx).Create(&match).Error; err != nil {
		return err
	}
	return nil
}
