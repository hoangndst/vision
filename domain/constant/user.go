package constant

import "errors"

var (
	ErrorUserNil           = errors.New("user is nil")
	ErrorUserNameEmpty     = errors.New("name is empty")
	ErrorUserUsernameEmpty = errors.New("username is empty")
	ErrorUserEmailEmpty    = errors.New("email is empty")
	ErrorUserPasswordEmpty = errors.New("password is empty")
	ErrorUserEmailInvalid  = errors.New("email is invalid")
	ErrorUserPasswordWeak  = errors.New("password is weak")
)

const (
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)
