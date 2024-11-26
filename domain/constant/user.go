package constant

import "errors"

var (
	ErrorUserNil           = errors.New("user is nil")
	ErrorUserNameEmpty     = errors.New("name is empty")
	ErrorUserUsernameEmpty = errors.New("username is empty")
	ErrorUserEmailEmpty    = errors.New("email is empty")
	ErrorUserPasswordEmpty = errors.New("password is empty")
)
