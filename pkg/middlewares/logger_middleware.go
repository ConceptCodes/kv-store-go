package middlewares

import (
	"bytes"
	"kv-store/pkg/models"
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context().Value("ctx").(*models.Request)

		log.Printf("[%s] %s %s", ctx.Id, r.Method, r.URL.Path)

		if r.Method == "POST" || r.Method == "PUT" {
			r.ParseForm()
			log.Printf("[%s] Request Body: %s", ctx.Id, r.Form)
		}

		next.ServeHTTP(w, r)

	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}
		ctx := r.Context().Value("ctx").(*models.Request)
		next.ServeHTTP(rw, r)
		log.Printf("[%s] Response Body: %s", ctx.Id, rw.body.String())
	})
}
