package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"kv-store/internal/constants"
	"kv-store/internal/helpers"
	"kv-store/internal/models"
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

			_user, err := helpers.ValidateToken(authToken)

			if err != nil {
				log.Error().Str("request_id", requestId).Msgf("Error: %s", err)
			}

			user = _user

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
