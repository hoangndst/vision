package models

import "errors"

var (
	ErrUserModelNil                = errors.New("user model is nil")
	ErrOrganizationModelNil        = errors.New("organization model is nil")
	ErrProjectModelNil             = errors.New("project model is nil")
	ErrFailedToConvertOrgToEntity  = errors.New("failed to convert organization to entity")
	ErrFailedToConvertUserToEntity = errors.New("failed to convert user to entity")
)

type OrganizationRole string

const (
	OrganizationRoleOwner OrganizationRole = "owner"
	OrganizationRoleAdmin OrganizationRole = "admin"
	OrganizationRoleUser  OrganizationRole = "user"
)
