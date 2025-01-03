package request

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/constant"
)

// CreateProjectRequest represents the create request structure for
// project.
type CreateProjectRequest struct {
	Name           string    `json:"name" binding:"required"`
	OrganizationID uuid.UUID `json:"organization_id" binding:"required"`
	Description    string    `json:"description"`
	Path           string    `json:"path"`
	Labels         []string  `json:"labels"`
}

// UpdateProjectRequest represents the update request structure for
// project.
type UpdateProjectRequest struct {
	// ID is the id of the project.
	ID                   uuid.UUID `json:"id" binding:"required"`
	CreateProjectRequest `json:",inline" yaml:",inline"`
}

func (payload *CreateProjectRequest) Validate() error {
	if payload.Name == "" {
		return constant.ErrProjectNameEmpty
	}
	if payload.OrganizationID == uuid.Nil {
		return constant.ErrorOrganizationNil
	}
	return nil
}

func (payload *UpdateProjectRequest) Validate() error {
	return payload.CreateProjectRequest.Validate()
}

func (payload *CreateProjectRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *UpdateProjectRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}
