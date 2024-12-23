package entity

import (
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/constant"
	"time"
)

type Project struct {
	ID                uuid.UUID     `yaml:"id" json:"id"`
	Name              string        `yaml:"name" json:"name"`
	Organization      *Organization `yaml:"organization" json:"organization"`
	Description       string        `yaml:"description,omitempty" json:"description,omitempty"`
	Path              string        `yaml:"path,omitempty" json:"path,omitempty"`
	Labels            []string      `yaml:"labels,omitempty" json:"labels,omitempty"`
	Owner             *User         `yaml:"owner,omitempty" json:"owner,omitempty"`
	CreationTimestamp time.Time     `yaml:"creationTimestamp,omitempty" json:"creationTimestamp,omitempty"`
	UpdateTimestamp   time.Time     `yaml:"updateTimestamp,omitempty" json:"updateTimestamp,omitempty"`
}

type ProjectFilter struct {
	OrgID uuid.UUID
	Name  string
}

func (p *Project) Validate() error {
	if p == nil {
		return constant.ErrProjectNil
	}
	if p.Name == "" {
		return constant.ErrProjectNameEmpty
	}
	if p.Organization == nil {
		return constant.ErrProjectOrganizationNil
	}
	if p.Owner == nil {
		return constant.ErrProjectOwnerNil
	}
	return nil
}
