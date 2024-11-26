package user

import (
	"github.com/go-chi/render"
	"github.com/hoangndst/vision/domain/request"
	"github.com/hoangndst/vision/server/handler"
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
//	@Router			/api/v1/users [post]
func (uh *UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logutil.GetLogger(ctx)
		logger.Info("Creating user")

		var requestPayload request.CreateUserRequest
		if err := requestPayload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		if err := requestPayload.Validate(); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		createdEntity, err := uh.userManager.CreateUser(ctx, requestPayload)
		handler.HandleResult(w, r, ctx, err, createdEntity)
	}
}
