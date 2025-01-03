package constant

import "errors"

var (
	ErrBlogNil          = errors.New("blog is nil")
	ErrBlogPathEmpty    = errors.New("path is empty")
	ErrBlogRawDataEmpty = errors.New("raw data is empty")
)
