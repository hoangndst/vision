package middleware

import (
	"context"
	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/server/manager/user"
	"github.com/hoangndst/vision/server/util/credentials"
	"net/http"
)

var (
	AuthContextKey   = &contextKey{"auth"}
	UserIDContextKey = &contextKey{"user"}
)

func InitAuthInfo() AuthInfo {
	return AuthInfo{
		AuthMethod: "",
		StatusCode: 0,
	}
}

type AuthMiddleware struct {
	userManager *user.UserManager
	logFilePath string
}

func NewAuthMiddleware(
	userManager *user.UserManager,
	logFilePath string,
) *AuthMiddleware {
	return &AuthMiddleware{
		userManager: userManager,
		logFilePath: logFilePath,
	}
}

func (m *AuthMiddleware) RequiredDev() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authInfo := r.Context().Value(AuthContextKey).(AuthInfo)
			if authInfo.AuthMethod == "" {
				http.Error(w, http.StatusText(authInfo.StatusCode), authInfo.StatusCode)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) RequiredPE() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authInfo := r.Context().Value(AuthContextKey).(AuthInfo)
			if authInfo.AuthMethod == "" {
				http.Error(w, http.StatusText(authInfo.StatusCode), authInfo.StatusCode)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func (m *AuthMiddleware) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authInfo, ok := r.Context().Value(AuthContextKey).(AuthInfo)
		if ok && authInfo.AuthMethod != "" {
			next.ServeHTTP(w, r)
			return
		}
		logger := GetMiddlewareLogger(r.Context(), m.logFilePath)
		authContext := InitAuthInfo()
		name, password, ok := r.BasicAuth()
		if !ok || name == "" || password == "" {
			logger.Info("basic auth is required")
			logger.Info("request denied")
			authContext.StatusCode = http.StatusUnauthorized
			ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		userEntity, err := m.userManager.GetUserWithPasswordByUsername(r.Context(), name)
		if err != nil {
			logger.Info("invalid user", "error", err)
			logger.Info("request denied")
			authContext.StatusCode = http.StatusUnauthorized
			ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ok, err = credentials.Compare(password, userEntity.Password)
		if err != nil {
			logger.Info("internal server error", "error", err)
			logger.Info("request denied")
			authContext.StatusCode = http.StatusInternalServerError
			ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		if !ok {
			logger.Info("invalid password")
			logger.Info("request denied")
			ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		authContext.AuthMethod = Basic

		ctx := context.WithValue(r.Context(), AuthContextKey, authContext)
		ctx = context.WithValue(ctx, UserIDContextKey, userEntity.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetMiddlewareLogger(ctx context.Context, logFile string) *httplog.Logger {
	if logger, ok := ctx.Value(APILoggerKey).(*httplog.Logger); ok {
		return logger
	}
	logger := InitLogger(logFile, "DefaultLogger")
	return logger
}

func GetUserID(ctx context.Context) uuid.UUID {
	if ctx == nil {
		return uuid.Nil
	}
	if userID, ok := ctx.Value(UserIDContextKey).(string); ok {
		return uuid.MustParse(userID)
	}
	return uuid.Nil
}
