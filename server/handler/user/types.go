package user

import (
	"github.com/google/uuid"
	usermanager "github.com/hoangndst/vision/server/manager/user"
)

func NewHandler(userManager *usermanager.UserManager) (*Handler, error) {
	return &Handler{
		userManager: userManager,
	}, nil
}

type Handler struct {
	userManager *usermanager.UserManager
}

type RequestParams struct {
	UserID uuid.UUID
}
