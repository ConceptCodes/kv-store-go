package middlewares

import (
	"kv-store/pkg/constants"
	"kv-store/pkg/helpers"
	"net/http"
)

func NotFound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.SendErrorResponse(w, "Not Found", constants.NotFound)
		return
	})
}
