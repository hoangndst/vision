package blog

import (
	"github.com/google/uuid"
	blogmanager "github.com/hoangndst/vision/server/manager/blog"
)

func NewHandler(blogManager *blogmanager.BlogManager) (*Handler, error) {
	return &Handler{
		blogManager: blogManager,
	}, nil
}

type Handler struct {
	blogManager *blogmanager.BlogManager
}

type RequestParams struct {
	BlogID   uuid.UUID
	BlogPath string
}
