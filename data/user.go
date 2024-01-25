package data

import (
	"context"
	"joust/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (m *UserRepo) Create(ctx context.Context, user *model.User) error {
	res := m.DB.WithContext(ctx).Create(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (m *UserRepo) ReadById(ctx context.Context, id uint) (*model.User, error) {

	return &model.User{}, nil
}

func (m *UserRepo) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (m *UserRepo) DeleteById(context.Context, uint) error {
	return nil
}
