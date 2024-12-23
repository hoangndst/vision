package models

import "github.com/hoangndst/vision/domain/entity"

type OrganizationModel struct {
	BaseModel
	Name   string `gorm:"index:unique_organization,unique"`
	Labels MultiString
	Users  []*UserModel `gorm:"many2many:user_organizations;"`
}

func (o *OrganizationModel) TableName() string {
	return "organizations"
}

func (o *OrganizationModel) ToEntity() (*entity.Organization, error) {
	if o == nil {
		return nil, ErrOrganizationModelNil
	}
	return &entity.Organization{
		ID:                o.ID,
		Name:              o.Name,
		Description:       o.Description,
		Labels:            o.Labels,
		CreationTimestamp: o.CreatedAt,
		UpdateTimestamp:   o.UpdatedAt,
	}, nil
}

func (o *OrganizationModel) FromEntity(entity *entity.Organization) error {
	if o == nil {
		return nil
	}
	o.ID = entity.ID
	o.Name = entity.Name
	o.Description = entity.Description
	o.Labels = entity.Labels
	o.CreatedAt = entity.CreationTimestamp
	o.UpdatedAt = entity.UpdateTimestamp
	return nil
}
