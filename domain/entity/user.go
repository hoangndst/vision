package entity

import (
	"github.com/hoangndst/vision/domain/constant"
	"time"
)

type User struct {
	ID                uint      `yaml:"id" json:"id"`
	Name              string    `yaml:"name" json:"name"`
	Username          string    `yaml:"username" json:"username"`
	Email             string    `yaml:"email" json:"email"`
	Password          string    `yaml:"password" json:"password"`
	CreationTimestamp time.Time `yaml:"creation_timestamp" json:"creation_timestamp"`
	UpdateTimestamp   time.Time `yaml:"update_timestamp" json:"update_timestamp"`
}

func (u *User) Validate() error {
	if u == nil {
		return constant.ErrorUserNil
	}
	if u.Name == "" {
		return constant.ErrorUserNameEmpty
	}
	if u.Username == "" {
		return constant.ErrorUserUsernameEmpty
	}
	if u.Email == "" {
		return constant.ErrorUserEmailEmpty
	}
	if u.Password == "" {
		return constant.ErrorUserPasswordEmpty
	}
	return nil
}
