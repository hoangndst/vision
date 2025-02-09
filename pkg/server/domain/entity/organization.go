package entity

import (
	"time"

	"github.com/hoangndst/vision/pkg/server/domain/constant"

	"github.com/google/uuid"
)

type Organization struct {
	ID                uuid.UUID `yaml:"id" json:"id"`
	Name              string    `yaml:"name" json:"name"`
	Description       string    `yaml:"description" json:"description"`
	Labels            []string  `yaml:"labels,omitempty" json:"labels,omitempty"`
	CreationTimestamp time.Time `yaml:"creation_timestamp" json:"creation_timestamp"`
	UpdateTimestamp   time.Time `yaml:"update_timestamp" json:"update_timestamp"`
}

func (o *Organization) Validate() error {
	if o == nil {
		return constant.ErrorOrganizationNil
	}
	if o.Name == "" {
		return constant.ErrorOrganizationNameEmpty
	}
	return nil
}
