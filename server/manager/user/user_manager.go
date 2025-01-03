package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/request"
	"github.com/hoangndst/vision/server/util/credentials"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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

func (u *UserManager) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
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

func (u *UserManager) GetUserPasswordByUsername(ctx context.Context, username string) (string, error) {
	password, err := u.userRepo.GetPasswordByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrGetNonExistentUser
		}
		return "", err
	}
	return password, nil
}

func (u *UserManager) GetUserWithPasswordByUsername(ctx context.Context, username string) (*entity.User, error) {
	userEntity, err := u.userRepo.GetWithPasswordByUsername(ctx, username)
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

	argon2iHash := credentials.NewDefaultArgon2idHash()
	hashedPassword, err := argon2iHash.HashPassword(createdEntity.Password, nil)
	if err != nil {
		return nil, err
	}
	createdEntity.Password = hashedPassword

	if err := u.userRepo.Create(ctx, &createdEntity); err != nil {
		return nil, err
	}
	return &createdEntity, nil
}

func (u *UserManager) UpdateUserByID(ctx context.Context, id uuid.UUID, requestPayload request.UpdateUserRequest) (*entity.User, error) {
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

func (m *UserManager) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	err := m.userRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGetNonExistentUser
		}
		return err
	}
	return nil
}
