package models

import (
	"github.com/hoangndst/vision/domain/entity"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string
	Username string `gorm:"index:unique_user,unique"`
	Email    string `gorm:"index:unique_user,unique"`
	Password string
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ToEntity() (*entity.User, error) {
	if u == nil {
		return nil, ErrUserModelNil
	}
	return &entity.User{
		ID:                u.ID,
		Name:              u.Name,
		Username:          u.Username,
		Email:             u.Email,
		Password:          u.Password,
		CreationTimestamp: u.CreatedAt,
		UpdateTimestamp:   u.UpdatedAt,
	}, nil
}

func (u *UserModel) FromEntity(entity *entity.User) error {
	if u == nil {
		return nil
	}
	u.ID = entity.ID
	u.Name = entity.Name
	u.Username = entity.Username
	u.Email = entity.Email
	u.Password = entity.Password
	u.CreatedAt = entity.CreationTimestamp
	u.UpdatedAt = entity.UpdateTimestamp
	return nil
}
