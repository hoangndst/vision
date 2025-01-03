package blog

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/request"
	"github.com/hoangndst/vision/server/handler"
	blogmanager "github.com/hoangndst/vision/server/manager/blog"
	logutil "github.com/hoangndst/vision/server/util/logging"
)

// @Id				createBlog
// @Summary		Create blog
// @Description	Create a new blog
// @Tags			blog
// @Accept			json
// @Produce		json
// @Param			blog	body		request.CreateBlogRequest			true	"Created blog"
// @Success		200		{object}	handler.Response{data=entity.Blog}	"Success"
// @Failure		400		{object}	error								"Bad Request"
// @Failure		401		{object}	error								"Unauthorized"
// @Failure		429		{object}	error								"Too Many Requests"
// @Failure		404		{object}	error								"Not Found"
// @Failure		500		{object}	error								"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs [post]
func (h *Handler) CreateBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Creating blog...")

		// Decode the request body into the payload.
		var requestPayload request.CreateBlogRequest
		if err := requestPayload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		err := requestPayload.Validate()
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		createdEntity, err := h.blogManager.CreateBlog(ctx, requestPayload)
		handler.HandleResult(w, r, ctx, err, createdEntity)
	}
}

// @Id				deleteBlog
// @Summary		Delete blog
// @Description	Delete specified blog by ID
// @Tags			blog
// @Produce		json
// @Param			id	path		string							true	"Blog ID"
// @Success		200	{object}	handler.Response{data=string}	"Success"
// @Failure		400	{object}	error							"Bad Request"
// @Failure		401	{object}	error							"Unauthorized"
// @Failure		429	{object}	error							"Too Many Requests"
// @Failure		404	{object}	error							"Not Found"
// @Failure		500	{object}	error							"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs/{id} [delete]
func (h *Handler) DeleteBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Deleting blog...", "blogID", params.BlogID)

		err = h.blogManager.DeleteBlogByID(ctx, params.BlogID)
		handler.HandleResult(w, r, ctx, err, "Deletion Success")
	}
}

// @Id				updateBlog
// @Summary		Update blog
// @Description	Update the specified blog
// @Tags			blog
// @Accept			json
// @Produce		json
// @Param			id		path		string								true	"Blog ID"
// @Param			blog	body		request.UpdateBlogRequest			true	"Updated blog"
// @Success		200		{object}	handler.Response{data=entity.Blog}	"Success"
// @Failure		400		{object}	error								"Bad Request"
// @Failure		401		{object}	error								"Unauthorized"
// @Failure		429		{object}	error								"Too Many Requests"
// @Failure		404		{object}	error								"Not Found"
// @Failure		500		{object}	error								"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs/{id} [put]
func (h *Handler) UpdateBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Updating blog...", "blogID", params.BlogID)

		// Decode the request body into the payload.
		var requestPayload request.UpdateBlogRequest
		if err := requestPayload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}

		updatedEntity, err := h.blogManager.UpdateBlogByID(ctx, params.BlogID, requestPayload)
		handler.HandleResult(w, r, ctx, err, updatedEntity)
	}
}

// @Id				getBlog
// @Summary		Get blog
// @Description	Get blog information by blog ID
// @Tags			blog
// @Produce		json
// @Param			id	path		string								true	"Blog ID"
// @Success		200	{object}	handler.Response{data=entity.Blog}	"Success"
// @Failure		400	{object}	error								"Bad Request"
// @Failure		401	{object}	error								"Unauthorized"
// @Failure		429	{object}	error								"Too Many Requests"
// @Failure		404	{object}	error								"Not Found"
// @Failure		500	{object}	error								"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs/{id} [get]
func (h *Handler) GetBlog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Getting blog...", "blogID", params.BlogID)

		existingEntity, err := h.blogManager.GetBlogByID(ctx, params.BlogID)
		handler.HandleResult(w, r, ctx, err, existingEntity)
	}
}

// @Id				listBlog
// @Summary		List blogs
// @Description	List all blogs
// @Tags			blog
// @Produce		json
// @Success		200	{object}	handler.Response{data=[]entity.Blog}	"Success"
// @Failure		400	{object}	error									"Bad Request"
// @Failure		401	{object}	error									"Unauthorized"
// @Failure		429	{object}	error									"Too Many Requests"
// @Failure		404	{object}	error									"Not Found"
// @Failure		500	{object}	error									"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs [get]
func (h *Handler) ListBlogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Listing blog...")

		blogEntities, err := h.blogManager.ListBlogs(ctx)
		handler.HandleResult(w, r, ctx, err, blogEntities)
	}
}

// @Id				syncBlogs
// @Summary		Sync blogs
// @Description	Sync blogs information from GitHub repository
// @Tags			blog
// @Produce		json
// @Success		200	{object}	handler.Response{data=string}	"Success"
// @Failure		400	{object}	error							"Bad Request"
// @Failure		401	{object}	error							"Unauthorized"
// @Failure		429	{object}	error							"Too Many Requests"
// @Failure		404	{object}	error							"Not Found"
// @Failure		500	{object}	error							"Internal Server Error"
// @Security		BasicAuth
// @Router			/api/v1/blogs/sync [post]
func (h *Handler) SyncBlogs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Syncing blogs...")

		err := h.blogManager.SyncBlogs(ctx)
		handler.HandleResult(w, r, ctx, err, "Sync Success")
	}
}

func requestHelper(r *http.Request) (context.Context, *httplog.Logger, *RequestParams, error) {
	ctx := r.Context()
	blogID := chi.URLParam(r, "blogID")
	blogPath := chi.URLParam(r, "blogPath")
	// convert string to uuid.UUID
	id, err := uuid.Parse(blogID)
	if err != nil {
		return nil, nil, nil, blogmanager.ErrInvalidBlogID
	}
	logger := logutil.GetLogger(ctx)
	params := RequestParams{
		BlogID:   id,
		BlogPath: blogPath,
	}
	return ctx, logger, &params, nil
}
