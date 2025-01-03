package models

import (
	"github.com/hoangndst/vision/domain/entity"
)

type BlogModel struct {
	ResourceModel
	Path    string `gorm:"index:unique_blog,unique"`
	RawData []byte `gorm:"type:bytea"`
}

// TableName returns the name of the database table that the struct is mapped to.
func (m *BlogModel) TableName() string {
	return "blogs"
}

// ToEntity converts the DO to an entity.
func (m *BlogModel) ToEntity() (*entity.Blog, error) {
	if m == nil {
		return nil, ErrBlogModelNil
	}

	return &entity.Blog{
		ID:                m.ID,
		Path:              m.Path,
		Description:       m.Description,
		RawData:           string(m.RawData),
		CreationTimestamp: m.CreatedAt,
		UpdateTimestamp:   m.UpdatedAt,
	}, nil
}

// FromEntity converts an entity to a DO.
func (m *BlogModel) FromEntity(e *entity.Blog) error {
	if m == nil {
		return nil
	}

	m.ID = e.ID
	m.Path = e.Path
	m.Description = e.Description
	m.RawData = []byte(e.RawData)
	m.CreatedAt = e.CreationTimestamp
	m.UpdatedAt = e.UpdateTimestamp
	return nil
}
