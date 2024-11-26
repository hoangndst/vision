package route

import (
	"context"
	"expvar"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	docs "github.com/hoangndst/vision/api/openapispec"
	"github.com/hoangndst/vision/models"
	"github.com/hoangndst/vision/server"
	"github.com/hoangndst/vision/server/handler/user"
	usermanager "github.com/hoangndst/vision/server/manager/user"
	"github.com/hoangndst/vision/server/middleware"
	logutil "github.com/hoangndst/vision/server/util/logging"
	httpswagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewRoute(config *server.Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware.TraceID)
	router.Use(middleware.UserID)
	router.Use(middleware.APILoggerMiddleware(config.LogFilePath))
	router.Use(middleware.DefaultLoggerMiddleware(config.LogFilePath))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Route("/api/v1", func(r chi.Router) {
		setupAPIV1(r, config)
	})

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Vision Backend API"
	docs.SwaggerInfo.Version = "v0.1.0"
	router.Get("/docs/*", httpswagger.Handler())

	router.Get("/configs", expvar.Handler().ServeHTTP)

	logger := logutil.GetLogger(context.TODO())
	logger.Info(fmt.Sprintf("Server is running on port %d", config.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		logger.Error(fmt.Sprintf("Failed to start server: %v", err))
		return nil, err
	}
	logger.Info("Server started")
	return router, nil
}

func setupAPIV1(r chi.Router, config *server.Config) {
	logger := logutil.GetLogger(context.TODO())
	logger.Info("Setting up API v1")

	if config.DB != nil && config.AutoMigrate {
		err := models.AutoMigrate(config.DB)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to auto migrate models: %v", err))
			return
		}
	}
	userRepo := models.NewUserRepository(config.DB)

	userManager := usermanager.NewUserManager(userRepo)

	userHandler, err := user.NewUserHandler(userManager)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create user handler: %v", err))
		return
	}
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser())
	})
}
