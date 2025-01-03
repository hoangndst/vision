package request

import (
	"net/http"
	"regexp"

	"github.com/hoangndst/vision/domain/constant"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (payload *CreateUserRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *UpdateUserRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *CreateUserRequest) Validate() error {
	if payload.Name == "" {
		return constant.ErrorUserNameEmpty
	}
	if payload.Username == "" {
		return constant.ErrorUserUsernameEmpty
	}
	if payload.Email == "" {
		return constant.ErrorUserEmailEmpty
	}
	if payload.Password == "" {
		return constant.ErrorUserPasswordEmpty
	}

	re := regexp.MustCompile(constant.EmailRegex)
	if !re.MatchString(payload.Email) {
		return constant.ErrorUserEmailInvalid
	}

	return nil
}

func (payload *UpdateUserRequest) Validate() error {
	if payload.Name == "" {
		return constant.ErrorUserNameEmpty
	}
	if payload.Email == "" {
		return constant.ErrorUserEmailEmpty
	}

	re := regexp.MustCompile(constant.EmailRegex)
	if !re.MatchString(payload.Email) {
		return constant.ErrorUserEmailInvalid
	}
	return nil
}
