package route

import (
	"context"
	"expvar"
	"fmt"
	"net/http"

	blogmodule "github.com/hoangndst/vision/pkg/modules/blog"
	"github.com/hoangndst/vision/pkg/server"
	"github.com/hoangndst/vision/pkg/server/handler/blog"
	"github.com/hoangndst/vision/pkg/server/handler/user"
	blogmanager "github.com/hoangndst/vision/pkg/server/manager/blog"
	usermanager "github.com/hoangndst/vision/pkg/server/manager/user"
	middleware2 "github.com/hoangndst/vision/pkg/server/middleware"
	"github.com/hoangndst/vision/pkg/server/models"
	logutil "github.com/hoangndst/vision/pkg/server/util/logging"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	docs "github.com/hoangndst/vision/api/openapispec"
	httpswagger "github.com/swaggo/http-swagger"
)

func NewRoute(config *server.Config) (*chi.Mux, error) {
	router := chi.NewRouter()

	router.Use(middleware2.TraceID)
	router.Use(middleware2.TraceUserID)
	router.Use(middleware2.Time)
	router.Use(middleware.Recoverer)
	router.Use(middleware2.APILoggerMiddleware(config.LogFilePath))
	router.Use(middleware2.DefaultLoggerMiddleware(config.LogFilePath))

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

	userRepo := models.NewUserRepository(config.DB)
	blogRepo := models.NewBlogRepository(config.DB)

	blogModule := blogmodule.NewClient(config.GithubToken)

	userManager := usermanager.NewUserManager(userRepo)
	blogManager := blogmanager.NewBlogManager(blogRepo, *blogModule)

	userHandler, err := user.NewHandler(userManager)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create user handler: %v", err))
		return
	}
	blogHandler, err := blog.NewHandler(blogManager)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create blog handler: %v", err))
		return
	}

	auth := middleware2.NewAuthMiddleware(userManager, config.LogFilePath)
	r.Use(auth.BasicAuth)

	r.Group(func(r chi.Router) {
		r.Use(auth.RequiredPE())
		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Get("/", userHandler.GetUser())
				r.Put("/", userHandler.UpdateUser())
				r.Delete("/", userHandler.DeleteUser())
			})
			r.Post("/", userHandler.CreateUser())
			r.Get("/", userHandler.ListUsers())
		})
		r.Route("/blogs", func(r chi.Router) {
			r.Route("/{blogID}", func(r chi.Router) {
				r.Get("/", blogHandler.GetBlog())
				r.Put("/", blogHandler.UpdateBlog())
				r.Delete("/", blogHandler.DeleteBlog())
			})
			r.Route("/path/{blogPath}", func(r chi.Router) {
				r.Get("/", blogHandler.GetBlogByPath())
			})
			r.Post("/", blogHandler.CreateBlog())
			r.Get("/", blogHandler.ListBlogs())
			r.Post("/sync", blogHandler.SyncBlogs())
			r.Get("/tags", blogHandler.GetTags())
		})
	})
}
