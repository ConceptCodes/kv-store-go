package middlewares

import (
	"bytes"
	"net/http"
	"kv-store/pkg/logger"
)

type responseWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := logger.GetLoggerWithContext(r.Context())

		log.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent())

		if r.Method == "POST" || r.Method == "PUT" {
			r.ParseForm()
			log.Info().Interface("data", r.Form).Msg("Request Data")
		}

		next.ServeHTTP(w, r)

	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}
		log := logger.GetLoggerWithContext(r.Context())

		log.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent()).
			Interface("response", rw.body.String()).
			Int("status_code", rw.statusCode)

		next.ServeHTTP(rw, r)
	})
}
