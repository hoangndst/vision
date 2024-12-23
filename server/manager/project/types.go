package project

import (
	"errors"
	"github.com/hoangndst/vision/domain/repository"
)

var (
	ErrGettingNonExistingProject  = errors.New("the project does not exist")
	ErrUpdatingNonExistingProject = errors.New("the project to update does not exist")
	ErrSourceNotFound             = errors.New("the specified source does not exist")
	ErrOrgNotFound                = errors.New("the specified org does not exist")
)

type ProjectManager struct {
	projectRepo      repository.ProjectRepository
	organizationRepo repository.OrganizationRepository
}

func NewProjectManager(projectRepo repository.ProjectRepository,
	organizationRepo repository.OrganizationRepository,
) *ProjectManager {
	return &ProjectManager{
		projectRepo:      projectRepo,
		organizationRepo: organizationRepo,
	}
}
