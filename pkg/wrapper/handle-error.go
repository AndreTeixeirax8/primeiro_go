package wrapper

import (
	"net/http"

	"github.com/go-chi/render"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandleError(endpointFunc EndpointFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, statusCode, err := endpointFunc(w, r)
		if err != nil {
			// You can customize error handling here if needed
			// For now, just use the statusCode returned by endpointFunc
			statusCode = http.StatusBadRequest
			render.Status(r, statusCode)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, statusCode)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	}
}