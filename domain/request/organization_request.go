package request

import "net/http"

// CreateOrganizationRequest represents the create request structure for
// organization.
type CreateOrganizationRequest struct {
	// Name is the name of the organization.
	Name string `json:"name" binding:"required"`
	// Description is a human-readable description of the organization.
	Description string `json:"description"`
	// Labels are custom labels associated with the organization.
	Labels []string `json:"labels"`
}

// UpdateOrganizationRequest represents the update request structure for
// organization.
type UpdateOrganizationRequest struct {
	CreateOrganizationRequest `json:",inline" yaml:",inline"`
}

func (payload *CreateOrganizationRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *UpdateOrganizationRequest) Decode(r *http.Request) error {
	return decode(r, payload)
}

func (payload *CreateOrganizationRequest) Validate() error {
	return nil
}

func (payload *UpdateOrganizationRequest) Validate() error {
	return payload.CreateOrganizationRequest.Validate()
}
