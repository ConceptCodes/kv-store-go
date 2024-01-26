package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"kv-store/pkg/constants"
	"kv-store/pkg/helpers"
	"kv-store/pkg/models"
)

func TraceRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestId := r.Header.Get(constants.TraceIdHeader)

		if requestId == "" {
			requestId = uuid.New().String()
		}

		w.Header().Add(constants.TraceIdHeader, requestId)

		authToken := r.Header.Get("Authorization")
		user := &models.UserModel{}

		if authToken != "" {
			authToken = strings.Replace(authToken, "Bearer ", "", 1)

			user, err := helpers.ValidateToken(authToken)

			if err != nil {
				log.Printf("Error: %s", err)
			}

			log.Printf("User Id: %s", user.ID)
		}

		model := &models.Request{
			Id:   requestId,
			User: *user,
		}

		ctx := context.WithValue(r.Context(), "ctx", model)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
