package constant

import "errors"

var (
	ErrProjectNil             = errors.New("project is nil")
	ErrProjectNameEmpty       = errors.New("name is empty")
	ErrProjectOrganizationNil = errors.New("organization is nil")
	ErrProjectOwnerNil        = errors.New("owner is nil")
)
