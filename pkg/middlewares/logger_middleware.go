package middlewares

import (
	"bytes"
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

		log.Printf("%s %s", r.Method, r.URL.Path)

		if r.Method == "POST" || r.Method == "PUT" {
			r.ParseForm()
			log.Printf("Request Body: %s", r.Form)
		}

		next.ServeHTTP(w, r)

	})
}

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}
		next.ServeHTTP(rw, r)
		log.Printf("Response Body: %s", rw.body.String())
	})
}
