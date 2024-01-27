package middlewares

import (
	"kv-store/internal/constants"
	"kv-store/internal/helpers"
	"net/http"
)

func NotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SendErrorResponse(w, "Not Found", constants.NotFound, nil)
		return
	})
}
