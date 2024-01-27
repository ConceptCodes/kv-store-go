package middlewares

import (
	"bytes"
	"kv-store/pkg/logger"
	"net/http"
	"time"
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

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log := logger.GetLoggerWithContext(r.Context())
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}

		start := time.Now()

		log.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Str("user_agent", r.UserAgent()).
			Dur("elapsed_ms", time.Since(start)).
			Int("status_code", rw.statusCode).
			Msgf("Incoming %s request", r.Method)

		next.ServeHTTP(w, r)

	})
}

func (lrw *responseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
