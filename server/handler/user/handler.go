package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/hoangndst/vision/domain/request"
	"github.com/hoangndst/vision/server/handler"
	usermanager "github.com/hoangndst/vision/server/manager/user"
	logutil "github.com/hoangndst/vision/server/util/logging"
	"net/http"
)

//	@Id				createUser
//	@Summary		Create user
//	@Description	Create a new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		request.CreateUserRequest	true	"Created user"
//	@Success		200		{object}	entity.User					"Success"
//	@Failure		400		{object}	error						"Bad Request"
//	@Failure		401		{object}	error						"Unauthorized"
//	@Failure		429		{object}	error						"Too Many Requests"
//	@Failure		404		{object}	error						"Not Found"
//	@Failure		500		{object}	error						"Internal Server Error"
//	@Security		BasicAuth
//	@Router			/api/v1/users [post]
func (h *Handler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Creating user...")

		// Decode the request body into the payload.
		var requestPayload request.CreateUserRequest
		if err := requestPayload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		err := requestPayload.Validate()
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		createdEntity, err := h.userManager.CreateUser(ctx, requestPayload)
		handler.HandleResult(w, r, ctx, err, createdEntity)
	}
}

//	@Id				deleteUser
//	@Summary		Delete user
//	@Description	Delete specified user by ID
//	@Tags			user
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	string	"Success"
//	@Failure		400	{object}	error	"Bad Request"
//	@Failure		401	{object}	error	"Unauthorized"
//	@Failure		429	{object}	error	"Too Many Requests"
//	@Failure		404	{object}	error	"Not Found"
//	@Failure		500	{object}	error	"Internal Server Error"
//	@Security		BasicAuth
//	@Router			/api/v1/users/{id} [delete]
func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Deleting user...", "userID", params.UserID)

		err = h.userManager.DeleteUserByID(ctx, params.UserID)
		handler.HandleResult(w, r, ctx, err, "Deletion Success")
	}
}

//	@Id				updateUser
//	@Summary		Update user
//	@Description	Update the specified user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"User ID"
//	@Param			user	body		request.UpdateUserRequest	true	"Updated user"
//	@Success		200		{object}	entity.User					"Success"
//	@Failure		400		{object}	error						"Bad Request"
//	@Failure		401		{object}	error						"Unauthorized"
//	@Failure		429		{object}	error						"Too Many Requests"
//	@Failure		404		{object}	error						"Not Found"
//	@Failure		500		{object}	error						"Internal Server Error"
//	@Security		BasicAuth
//	@Router			/api/v1/users/{id} [put]
func (h *Handler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Updating user...", "userID", params.UserID)

		// Decode the request body into the payload.
		var requestPayload request.UpdateUserRequest
		if err := requestPayload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}

		updatedEntity, err := h.userManager.UpdateUserByID(ctx, params.UserID, requestPayload)
		handler.HandleResult(w, r, ctx, err, updatedEntity)
	}
}

//	@Id				getUser
//	@Summary		Get user
//	@Description	Get user information by user ID
//	@Tags			user
//	@Produce		json
//	@Param			id	path		string		true	"User ID"
//	@Success		200	{object}	entity.User	"Success"
//	@Failure		400	{object}	error		"Bad Request"
//	@Failure		401	{object}	error		"Unauthorized"
//	@Failure		429	{object}	error		"Too Many Requests"
//	@Failure		404	{object}	error		"Not Found"
//	@Failure		500	{object}	error		"Internal Server Error"
//	@Security		BasicAuth
//	@Router			/api/v1/users/{id} [get]
func (h *Handler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx, logger, params, err := requestHelper(r)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		logger.Info("Getting user...", "userID", params.UserID)

		existingEntity, err := h.userManager.GetUserByID(ctx, params.UserID)
		handler.HandleResult(w, r, ctx, err, existingEntity)
	}
}

//	@Id				listUser
//	@Summary		List users
//	@Description	List all users
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	entity.User	"Success"
//	@Failure		400	{object}	error		"Bad Request"
//	@Failure		401	{object}	error		"Unauthorized"
//	@Failure		429	{object}	error		"Too Many Requests"
//	@Failure		404	{object}	error		"Not Found"
//	@Failure		500	{object}	error		"Internal Server Error"
//	@Security		BasicAuth
//	@Router			/api/v1/users [get]
func (h *Handler) ListUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Getting stuff from context
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Listing user...")

		userEntities, err := h.userManager.ListUsers(ctx)
		handler.HandleResult(w, r, ctx, err, userEntities)
	}
}

func requestHelper(r *http.Request) (context.Context, *httplog.Logger, *RequestParams, error) {
	ctx := r.Context()
	userID := chi.URLParam(r, "userID")
	// convert string to uuid.UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, nil, nil, usermanager.ErrInvalidUserID
	}
	logger := logutil.GetLogger(ctx)
	params := RequestParams{
		UserID: id,
	}
	return ctx, logger, &params, nil
}
