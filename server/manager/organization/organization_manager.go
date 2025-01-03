package organization

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/request"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (m *OrganizationManager) ListOrganizations(ctx context.Context) ([]*entity.Organization, error) {
	organizationEntities, err := m.organizationRepo.List(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingOrganization
		}
		return nil, err
	}
	return organizationEntities, nil
}

func (m *OrganizationManager) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	existingEntity, err := m.organizationRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingOrganization
		}
		return nil, err
	}
	return existingEntity, nil
}

func (m *OrganizationManager) DeleteOrganizationByID(ctx context.Context, id uuid.UUID) error {
	err := m.organizationRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGettingNonExistingOrganization
		}
		return err
	}
	return nil
}

func (m *OrganizationManager) UpdateOrganizationByID(ctx context.Context, id uuid.UUID, requestPayload request.UpdateOrganizationRequest) (*entity.Organization, error) {
	// Convert request payload to domain model
	var requestEntity entity.Organization
	if err := copier.Copy(&requestEntity, &requestPayload); err != nil {
		return nil, err
	}

	// Get the existing organization by id
	updatedEntity, err := m.organizationRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUpdatingNonExistingOrganization
		}
		return nil, err
	}

	// Overwrite non-zero values in request entity to existing entity
	copier.CopyWithOption(updatedEntity, requestEntity, copier.Option{IgnoreEmpty: true})

	// Update organization with repository
	err = m.organizationRepo.Update(ctx, updatedEntity)
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (m *OrganizationManager) CreateOrganization(ctx context.Context, requestPayload request.CreateOrganizationRequest) (*entity.Organization, error) {
	// Convert request payload to domain model
	var createdEntity entity.Organization
	if err := copier.Copy(&createdEntity, &requestPayload); err != nil {
		return nil, err
	}

	// Create organization with repository
	err := m.organizationRepo.Create(ctx, &createdEntity)
	if err != nil {
		return nil, err
	}
	return &createdEntity, nil
}
