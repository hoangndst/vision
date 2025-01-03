package models

import (
	"context"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/repository"
	"gorm.io/gorm"
)

var _ repository.BlogRepository = &blogRepository{}

// blogRepository is a repository that stores blogs in a gorm database.
type blogRepository struct {
	// db is the underlying gorm database where blogs are stored.
	db *gorm.DB
}

// NewBlogRepository creates a new blog repository.
func NewBlogRepository(db *gorm.DB) repository.BlogRepository {
	return &blogRepository{db: db}
}

// Create saves a blog to the repository.
func (r *blogRepository) Create(ctx context.Context, dataEntity *entity.Blog) error {
	err := dataEntity.Validate()
	if err != nil {
		return err
	}

	// Map the data from Entity to DO
	var dataModel BlogModel
	err = dataModel.FromEntity(dataEntity)
	if err != nil {
		return err
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

// Delete removes a blog from the repository.
func (r *blogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var dataModel BlogModel
		err := tx.WithContext(ctx).First(&dataModel, id).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Delete(&dataModel).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// Update updates an existing blog in the repository.
func (r *blogRepository) Update(ctx context.Context, dataEntity *entity.Blog) error {
	// Map the data from Entity to DO
	var dataModel BlogModel
	err := dataModel.FromEntity(dataEntity)
	if err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).Model(&dataModel).Updates(&dataModel).Error
		if err != nil {
			return err
		}

		return nil
	})
}

// Get retrieves a blog by its ID.
func (r *blogRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Blog, error) {
	var dataModel BlogModel
	err := r.db.WithContext(ctx).First(&dataModel, id).Error
	if err != nil {
		return nil, err
	}

	dataEntity, err := dataModel.ToEntity()
	if err != nil {
		return nil, err
	}

	return dataEntity, nil
}

// GetByName retrieves a blog by its path.
func (r *blogRepository) GetByName(ctx context.Context, path string) (*entity.Blog, error) {
	var dataModel BlogModel
	err := r.db.WithContext(ctx).Where("path = ?", path).First(&dataModel).Error
	if err != nil {
		return nil, err
	}

	dataEntity, err := dataModel.ToEntity()
	if err != nil {
		return nil, err
	}

	return dataEntity, nil
}

// List retrieves all blogs from the repository.
func (r *blogRepository) List(ctx context.Context) ([]*entity.Blog, error) {
	var dataModels []BlogModel
	err := r.db.WithContext(ctx).Find(&dataModels).Error
	if err != nil {
		return nil, err
	}

	var dataEntities []*entity.Blog
	for _, dataModel := range dataModels {
		dataEntity, err := dataModel.ToEntity()
		if err != nil {
			return nil, err
		}
		dataEntities = append(dataEntities, dataEntity)
	}

	return dataEntities, nil
}

// GetByPath retrieves a blog by its path.
func (r *blogRepository) GetByPath(ctx context.Context, path string) (*entity.Blog, error) {
	var dataModel BlogModel
	err := r.db.WithContext(ctx).Where("path = ?", path).First(&dataModel).Error
	if err != nil {
		return nil, err
	}

	dataEntity, err := dataModel.ToEntity()
	if err != nil {
		return nil, err
	}

	return dataEntity, nil
}
