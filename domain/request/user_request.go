package request

import (
	"net/http"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	// ID is the id of the user.
	ID                uint `json:"id"`
	CreateUserRequest `json:",inline" yaml:",inline"`
}

func (payload *CreateUserRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *UpdateUserRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *CreateUserRequest) Validate() error {
	// TODO: Implement validation logic
	return nil
}

func (payload *UpdateUserRequest) Validate() error {
	// TODO: Implement validation logic
	return nil
}
