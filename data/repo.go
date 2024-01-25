package data

import (
	"context"
)

type Repo[T any] interface {
	Create(context.Context, *T) error
	ReadById(context.Context, uint) (*T, error)
	Update(context.Context, *T) error
	DeleteById(context.Context, uint) error
}
