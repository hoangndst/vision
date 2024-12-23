package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetPasswordByUsername(ctx context.Context, username string) (string, error)
	GetWithPasswordByUsername(ctx context.Context, username string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
}

type OrganizationRepository interface {
	Create(ctx context.Context, organization *entity.Organization) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, organization *entity.Organization) error
	Get(ctx context.Context, id uuid.UUID) (*entity.Organization, error)
	GetByName(ctx context.Context, name string) (*entity.Organization, error)
	GetUsers(ctx context.Context, id uuid.UUID) ([]*entity.User, error)
	List(ctx context.Context) ([]*entity.Organization, error)
}

type ProjectRepository interface {
	Create(ctx context.Context, project *entity.Project) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, project *entity.Project) error
	Get(ctx context.Context, id uuid.UUID) (*entity.Project, error)
	GetByName(ctx context.Context, name string) (*entity.Project, error)
	List(ctx context.Context, filter *entity.ProjectFilter) ([]*entity.Project, error)
}
