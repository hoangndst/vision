package models

import (
	"gorm.io/gorm"
	"time"
)

type UserOrganizationModel struct {
	UserID         string `gorm:"primaryKey"`
	OrganizationID string `gorm:"primaryKey"`
	Role           OrganizationRole
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (uo *UserOrganizationModel) TableName() string {
	return "user_organizations"
}
