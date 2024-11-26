package repository

import (
	"context"
	"github.com/hoangndst/vision/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id uint) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
}
