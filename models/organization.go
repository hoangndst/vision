package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/server/middleware"

	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/repository"

	"gorm.io/gorm"
)

// The organizationRepository type implements the repository.OrganizationRepository interface.
// If the organizationRepository type does not implement all the methods of the interface,
// the compiler will produce an error.
var _ repository.OrganizationRepository = &organizationRepository{}

// organizationRepository is a repository that stores organizations in a gorm database.
type organizationRepository struct {
	// db is the underlying gorm database where organizations are stored.
	db *gorm.DB
}

// NewOrganizationRepository creates a new organization repository.
func NewOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return &organizationRepository{db: db}
}

// Create saves a organization to the repository.
func (r *organizationRepository) Create(ctx context.Context, dataEntity *entity.Organization) error {
	err := dataEntity.Validate()
	if err != nil {
		return err
	}

	// Map the data from Entity to DO
	var dataModel OrganizationModel
	err = dataModel.FromEntity(dataEntity)
	if err != nil {
		return err
	}

	userID := middleware.GetUserID(ctx)
	if userID != uuid.Nil {
		var userModel UserModel
		err = r.db.WithContext(ctx).First(&userModel, userID).Error
		if err != nil {
			return err
		}
		dataModel.Users = []*UserModel{&userModel}
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create new record in the store
		err = tx.WithContext(ctx).Create(&dataModel).Error
		if err != nil {
			return err
		}

		dataEntity.ID = dataModel.ID

		return nil
	})
}

// Delete removes a organization from the repository.
func (r *organizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var dataModel OrganizationModel
		err := tx.WithContext(ctx).First(&dataModel, id).Error
		if err != nil {
			return err
		}

		return tx.WithContext(ctx).Unscoped().Delete(&dataModel).Error
	})
}

// Update updates an existing organization in the repository.
func (r *organizationRepository) Update(ctx context.Context, dataEntity *entity.Organization) error {
	// Map the data from Entity to DO
	var dataModel OrganizationModel
	err := dataModel.FromEntity(dataEntity)
	if err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Updates(&dataModel).Error
	if err != nil {
		return err
	}

	return nil
}

// Get retrieves a organization by its ID.
func (r *organizationRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Organization, error) {
	var dataModel OrganizationModel
	err := r.db.WithContext(ctx).First(&dataModel, id).Error
	if err != nil {
		return nil, err
	}

	return dataModel.ToEntity()
}

// GetByName retrieves a organization by its name.
func (r *organizationRepository) GetByName(ctx context.Context, name string) (*entity.Organization, error) {
	var dataModel OrganizationModel
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&dataModel).Error
	if err != nil {
		return nil, err
	}
	return dataModel.ToEntity()
}

// GetUsers retrieves all users of an organization.
func (r *organizationRepository) GetUsers(ctx context.Context, id uuid.UUID) ([]*entity.User, error) {
	var dataModel OrganizationModel
	err := r.db.WithContext(ctx).Preload("Users").First(&dataModel, id).Error
	if err != nil {
		return nil, err
	}

	userEntityList := make([]*entity.User, 0)
	for _, user := range dataModel.Users {
		userEntity, err := user.ToEntity()
		if err != nil {
			return nil, err
		}
		userEntityList = append(userEntityList, userEntity)
	}

	return userEntityList, nil
}

// List retrieves all organizations.
func (r *organizationRepository) List(ctx context.Context) ([]*entity.Organization, error) {
	var dataModel []OrganizationModel
	organizationEntityList := make([]*entity.Organization, 0)
	result := r.db.WithContext(ctx).Find(&dataModel)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, organization := range dataModel {
		organizationEntity, err := organization.ToEntity()
		if err != nil {
			return nil, err
		}
		organizationEntityList = append(organizationEntityList, organizationEntity)
	}
	return organizationEntityList, nil
}
