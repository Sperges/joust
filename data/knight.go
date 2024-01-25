package data

import (
	"context"
	"joust/model"

	"gorm.io/gorm"
)

type KnightRepo struct {
	DB *gorm.DB
}

func (r *KnightRepo) RandomFirstName(ctx context.Context) (string, error) {
	return RandomFromTable(ctx, r.DB, "first_names")
}

func (r *KnightRepo) RandomLastName(ctx context.Context) (string, error) {
	return RandomFromTable(ctx, r.DB, "last_names")
}

func (r *KnightRepo) RandomRegion(ctx context.Context) (string, error) {
	return RandomFromTable(ctx, r.DB, "regions")
}

func (r *KnightRepo) Create(ctx context.Context, knight *model.Knight) error {
	if err := r.DB.WithContext(ctx).Create(knight).Error; err != nil {
		return err
	}
	return nil
}

func (r *KnightRepo) ReadById(ctx context.Context, id uint) (*model.Knight, error) {
	knight := &model.Knight{}
	res := r.DB.WithContext(ctx).Model(knight).Preload("KnightStats").First(knight, id)
	if res.Error != nil {
		return &model.Knight{}, res.Error
	}
	return knight, nil
}

func (r *KnightRepo) Update(context.Context, *model.Knight) error {
	return nil
}

func (r *KnightRepo) DeleteById(context.Context, uint) error {
	return nil
}
