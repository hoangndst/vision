package middleware

type contextKey struct {
	name string
}
type AuthMethod string

const (
	Basic  AuthMethod = "BasicAuth"
	APIKey AuthMethod = "APIKeyAuth"
)

type AuthInfo struct {
	AuthMethod AuthMethod
	StatusCode int
}
