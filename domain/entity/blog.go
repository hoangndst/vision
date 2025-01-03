package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/constant"
)

type Blog struct {
	ID                uuid.UUID `yaml:"id" json:"id"`
	Path              string    `yaml:"path" json:"path"`
	Description       string    `yaml:"description,omitempty" json:"description,omitempty"`
	RawData           string    `yaml:"raw_data" json:"raw_data"`
	CreationTimestamp time.Time `yaml:"creation_timestamp" json:"creation_timestamp"`
	UpdateTimestamp   time.Time `yaml:"update_timestamp" json:"update_timestamp"`
}

func (b *Blog) Validate() error {
	if b == nil {
		return constant.ErrBlogNil
	}
	if b.Path == "" {
		return constant.ErrBlogPathEmpty
	}
	if b.RawData == "" {
		return constant.ErrBlogRawDataEmpty
	}
	return nil
}
