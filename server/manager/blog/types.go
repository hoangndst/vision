package blog

import (
	"errors"

	"github.com/hoangndst/vision/domain/repository"
	"github.com/hoangndst/vision/modules/blog"
)

var (
	ErrGettingNonExistingBlog  = errors.New("blog does not exist")
	ErrUpdatingNonExistingBlog = errors.New("blog to update does not exist")
	ErrInvalidBlogID           = errors.New("invalid blog id")
)

type BlogManager struct {
	blogRepo   repository.BlogRepository
	blogModule blog.Client
}

func NewBlogManager(blogRepo repository.BlogRepository, blogModule blog.Client) *BlogManager {
	return &BlogManager{
		blogRepo:   blogRepo,
		blogModule: blogModule,
	}
}
