package user

import (
	usermanager "github.com/hoangndst/vision/server/manager/user"
)

func NewUserHandler(userManager *usermanager.UserManager) (*UserHandler, error) {
	return &UserHandler{
		userManager: userManager,
	}, nil
}

type UserHandler struct {
	userManager *usermanager.UserManager
}
