package models

import (
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
)

type ProjectModel struct {
	ResourceModel
	Name           string `gorm:"index:unique_project,unique"`
	OrganizationID uuid.UUID
	Organization   *OrganizationModel `gorm:"foreignKey:OrganizationID;references:ID"`
	Path           string             `gorm:"index:unique_project,unique"`
	Labels         MultiString
}

// The TableName method returns the name of the database table that the struct is mapped to.
func (m *ProjectModel) TableName() string {
	return "project"
}

// ToEntity converts the DO to an entity.
func (m *ProjectModel) ToEntity() (*entity.Project, error) {
	if m == nil {
		return nil, ErrProjectModelNil
	}

	var err error
	var userEntity *entity.User
	var organizationEntity *entity.Organization

	if m.CreatedBy != nil {
		userEntity, err = m.CreatedBy.ToEntity()
		if err != nil {
			return nil, ErrFailedToConvertUserToEntity
		}
	}

	if m.Organization != nil {
		organizationEntity, err = m.Organization.ToEntity()
		if err != nil {
			return nil, ErrFailedToConvertOrgToEntity
		}
	}

	return &entity.Project{
		ID:                m.ID,
		Name:              m.Name,
		Organization:      organizationEntity,
		Owner:             userEntity,
		Path:              m.Path,
		Description:       m.Description,
		Labels:            []string(m.Labels),
		CreationTimestamp: m.CreatedAt,
		UpdateTimestamp:   m.UpdatedAt,
	}, nil
}

// FromEntity converts an entity to a DO.
func (m *ProjectModel) FromEntity(e *entity.Project) error {
	if m == nil {
		return ErrProjectModelNil
	}

	m.ID = e.ID
	m.Name = e.Name
	m.Description = e.Description
	m.Path = e.Path
	m.Labels = MultiString(e.Labels)
	m.CreatedAt = e.CreationTimestamp
	m.UpdatedAt = e.UpdateTimestamp
	if e.Organization != nil {
		m.OrganizationID = e.Organization.ID
		m.Organization.FromEntity(e.Organization)
	}
	if e.Owner != nil {
		m.CreatedByID = e.Owner.ID
		m.CreatedBy.FromEntity(e.Owner)
	}
	return nil
}
