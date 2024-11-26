package user

import (
	"context"
	"errors"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/request"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

func (u *UserManager) ListUsers(ctx context.Context) ([]*entity.User, error) {
	userEntities, err := u.userRepo.List(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetNonExistentUser
		}
		return nil, err
	}
	return userEntities, nil
}

func (u *UserManager) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	userEntity, err := u.userRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetNonExistentUser
		}
		return nil, err
	}
	return userEntity, nil
}

func (u *UserManager) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	userEntity, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetNonExistentUser
		}
		return nil, err
	}
	return userEntity, nil
}

func (u *UserManager) CreateUser(ctx context.Context, requestPayload request.CreateUserRequest) (*entity.User, error) {
	var createdEntity entity.User
	if err := copier.Copy(&createdEntity, &requestPayload); err != nil {
		return nil, err
	}
	createdEntity.CreationTimestamp = time.Now()
	createdEntity.UpdateTimestamp = time.Now()

	if err := u.userRepo.Create(ctx, &createdEntity); err != nil {
		return nil, err
	}
	return &createdEntity, nil
}

func (u *UserManager) UpdateUser(ctx context.Context, id uint, requestPayload request.UpdateUserRequest) (*entity.User, error) {
	var requestEntity entity.User
	if err := copier.Copy(&requestEntity, &requestPayload); err != nil {
		return nil, err
	}
	updatedEntity, err := u.userRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUpdateNonExistentUser
		}
		return nil, err
	}
	copier.CopyWithOption(updatedEntity, requestEntity, copier.Option{IgnoreEmpty: true})
	if err = u.userRepo.Update(ctx, updatedEntity); err != nil {
		return nil, err
	}
	return updatedEntity, nil
}
