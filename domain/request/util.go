package request

import (
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

func decode(r *http.Request, payload interface{}) error {
	contentType := render.GetRequestContentType(r)
	switch contentType {
	case render.ContentTypeJSON:
		if err := render.DecodeJSON(r.Body, payload); err != nil {
			return err
		}
	default:
		return errors.New("unsupported content type")
	}
	return nil
}
