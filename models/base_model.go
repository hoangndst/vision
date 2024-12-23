package models

import (
	"github.com/google/uuid"
	"github.com/hoangndst/vision/server/middleware"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Description string
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}

type ResourceModel struct {
	BaseModel
	CreatedByID uuid.UUID
	CreatedBy   *UserModel `gorm:"foreignKey:CreatedByID;references:ID"`
}

func (r *ResourceModel) BeforeCreate(tx *gorm.DB) error {
	createdByID := middleware.GetUserID(tx.Statement.Context)
	r.CreatedByID = createdByID
	return nil
}
