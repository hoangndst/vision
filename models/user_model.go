package models

import (
	"github.com/hoangndst/vision/domain/entity"
)

type UserModel struct {
	BaseModel
	Name          string
	Username      string               `gorm:"index:unique_user,unique"`
	Email         string               `gorm:"index:unique_user,unique"`
	Password      string               `json:"-"`
	Organizations []*OrganizationModel `gorm:"many2many:user_organizations;"`
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
		Description:       u.Description,
		Username:          u.Username,
		Email:             u.Email,
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
	u.Description = entity.Description
	u.Username = entity.Username
	u.Email = entity.Email
	u.Password = entity.Password
	u.CreatedAt = entity.CreationTimestamp
	u.UpdatedAt = entity.UpdateTimestamp
	return nil
}
