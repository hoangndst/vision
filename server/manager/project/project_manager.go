package project

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/request"
	logutil "github.com/hoangndst/vision/server/util/logging"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (m *ProjectManager) ListProjects(ctx context.Context, filter *entity.ProjectFilter) ([]*entity.Project, error) {
	projectEntities, err := m.projectRepo.List(ctx, filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingProject
		}
		return nil, err
	}
	return projectEntities, nil
}

func (m *ProjectManager) GetProjectByID(ctx context.Context, id uuid.UUID) (*entity.Project, error) {
	existingEntity, err := m.projectRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingProject
		}
		return nil, err
	}
	return existingEntity, nil
}

func (m *ProjectManager) DeleteProjectByID(ctx context.Context, id uuid.UUID) error {
	err := m.projectRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGettingNonExistingProject
		}
		return err
	}
	return nil
}

func (m *ProjectManager) UpdateProjectByID(ctx context.Context, id uuid.UUID, requestPayload request.UpdateProjectRequest) (*entity.Project, error) {
	// Convert request payload to domain model
	var requestEntity entity.Project
	if err := copier.Copy(&requestEntity, &requestPayload); err != nil {
		return nil, err
	}

	// Get the existing project by id
	updatedEntity, err := m.projectRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUpdatingNonExistingProject
		}
		return nil, err
	}

	// Get organization by id
	if requestPayload.OrganizationID == uuid.Nil {
		requestEntity.Organization = updatedEntity.Organization
	} else {
		// If orgID is passed in, get org by id and update the project organization
		organizationEntity, err := m.organizationRepo.Get(ctx, requestPayload.OrganizationID)
		if err != nil {
			return nil, err
		}
		requestEntity.Organization = organizationEntity
	}

	// Overwrite non-zero values in request entity to existing entity
	copier.CopyWithOption(updatedEntity, requestEntity, copier.Option{IgnoreEmpty: true})

	// Update project with repository
	err = m.projectRepo.Update(ctx, updatedEntity)
	if err != nil {
		return nil, err
	}
	return updatedEntity, nil
}

func (m *ProjectManager) CreateProject(ctx context.Context, requestPayload request.CreateProjectRequest) (*entity.Project, error) {
	logger := logutil.GetLogger(ctx)
	// Convert request payload to domain model
	var createdEntity entity.Project
	if err := copier.Copy(&createdEntity, &requestPayload); err != nil {
		return nil, err
	}

	// If orgID is passed in, get org by id
	if requestPayload.OrganizationID != uuid.Nil {
		logger.Info("Organization ID found in the request. Using the organization ID...", "organizationID", requestPayload.OrganizationID)
		organizationEntity, err := m.organizationRepo.Get(ctx, requestPayload.OrganizationID)
		if err != nil {
			return nil, err
		}
		createdEntity.Organization = organizationEntity
	}
	// Create project with repository
	err := m.projectRepo.Create(ctx, &createdEntity)
	if err != nil {
		return nil, err
	}
	return &createdEntity, nil
}
