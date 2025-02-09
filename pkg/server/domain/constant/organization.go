package constant

import "errors"

var (
	ErrorOrganizationNil       = errors.New("organization is nil")
	ErrorOrganizationNameEmpty = errors.New("name is empty")
)
