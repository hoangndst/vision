package request

import "net/http"

// CreateBlogRequest represents the create request structure for blog.
type CreateBlogRequest struct {
	// Path is the path of the blog.
	Path string `json:"path" binding:"required"`
	// RawData is the raw data of the blog.
	RawData string `json:"raw_data" binding:"required"`
	// Description is a human-readable description of the blog.
	Description string `json:"description"`
}

// UpdateBlogRequest represents the update request structure for blog.
type UpdateBlogRequest struct {
	CreateBlogRequest `json:",inline" yaml:",inline"`
}

func (payload *CreateBlogRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *UpdateBlogRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *CreateBlogRequest) Validate() error {
	return nil
}

func (payload *UpdateBlogRequest) Validate() error {
	return payload.CreateBlogRequest.Validate()
}
