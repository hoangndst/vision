package organization

import (
	"errors"

	"github.com/hoangndst/vision/domain/repository"
)

var (
	ErrGettingNonExistingOrganization  = errors.New("the organization does not exist")
	ErrUpdatingNonExistingOrganization = errors.New("the organization to update does not exist")
)

type OrganizationManager struct {
	organizationRepo repository.OrganizationRepository
}

func NewOrganizationManager(organizationRepo repository.OrganizationRepository) *OrganizationManager {
	return &OrganizationManager{
		organizationRepo: organizationRepo,
	}
}
