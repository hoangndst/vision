package blog

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/entity"
	"github.com/hoangndst/vision/domain/request"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func (m *BlogManager) ListBlogs(ctx context.Context) ([]*entity.Blog, error) {
	blogEntities, err := m.blogRepo.List(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingBlog
		}
		return nil, err
	}
	return blogEntities, nil
}

func (m *BlogManager) GetBlogByID(ctx context.Context, id uuid.UUID) (*entity.Blog, error) {
	existingEntity, err := m.blogRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingBlog
		}
		return nil, err
	}
	return existingEntity, nil
}

func (m *BlogManager) GetBlogByPath(ctx context.Context, path string) (*entity.Blog, error) {
	existingEntity, err := m.blogRepo.GetByPath(ctx, path)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGettingNonExistingBlog
		}
		return nil, err
	}
	return existingEntity, nil
}

func (m *BlogManager) DeleteBlogByID(ctx context.Context, id uuid.UUID) error {
	err := m.blogRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrGettingNonExistingBlog
		}
	}
	return nil
}

func (m *BlogManager) UpdateBlogByID(ctx context.Context, id uuid.UUID, requestPayload request.UpdateBlogRequest) (*entity.Blog, error) {
	// Convert request payload to domain model
	var requestEntity entity.Blog
	if err := copier.Copy(&requestEntity, &requestPayload); err != nil {
		return nil, err
	}

	// Get the existing blog by id
	updatedEntity, err := m.blogRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUpdatingNonExistingBlog
		}
		return nil, err
	}

	// Overwrite non-zero values in request entity to existing entity
	copier.CopyWithOption(updatedEntity, requestEntity, copier.Option{IgnoreEmpty: true})

	// Update blog with repository
	err = m.blogRepo.Update(ctx, updatedEntity)
	if err != nil {
		return nil, err
	}

	return updatedEntity, nil
}

func (m *BlogManager) CreateBlog(ctx context.Context, requestPayload request.CreateBlogRequest) (*entity.Blog, error) {
	// Convert request payload to domain model
	var requestEntity entity.Blog
	if err := copier.Copy(&requestEntity, &requestPayload); err != nil {
		return nil, err
	}

	// Create blog with repository
	err := m.blogRepo.Create(ctx, &requestEntity)
	if err != nil {
		return nil, err
	}

	return &requestEntity, nil
}

func (m *BlogManager) SyncBlogs(ctx context.Context) error {
	blogs, err := m.blogModule.GetAllBlogs(ctx)
	if err != nil {
		return err
	}

	for _, b := range blogs {
		blogEntity, err := m.GetBlogByPath(ctx, b.Path)
		if err != nil || errors.Is(err, ErrGettingNonExistingBlog) {
			requestPayload := request.CreateBlogRequest{
				Path:        b.Path,
				RawData:     b.RawData,
				Description: "Blog synced from external source",
			}
			_, err := m.CreateBlog(ctx, requestPayload)
			if err != nil {
				return err
			}
		} else {
			requestPayload := request.UpdateBlogRequest{
				CreateBlogRequest: request.CreateBlogRequest{
					Path:        b.Path,
					RawData:     b.RawData,
					Description: "Blog synced from external source",
				},
			}
			_, err := m.UpdateBlogByID(ctx, blogEntity.ID, requestPayload)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
