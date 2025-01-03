package user

import (
	"errors"

	"github.com/hoangndst/vision/domain/repository"
)

var (
	ErrGetNonExistentUser    = errors.New("the user does not exist")
	ErrUpdateNonExistentUser = errors.New("the user does not exist")
	ErrInvalidUserID         = errors.New("the user ID should be a uuid")
)

type UserManager struct {
	userRepo repository.UserRepository
}

func NewUserManager(userRepo repository.UserRepository) *UserManager {
	return &UserManager{
		userRepo: userRepo,
	}
}
