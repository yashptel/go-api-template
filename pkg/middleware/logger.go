package middleware

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type (
	// struct for holding response details
	responseData struct {
		status int
		body   []byte
	}

	// our http.ResponseWriter implementation
	loggingResponseWriter struct {
		http.ResponseWriter // compose original http.ResponseWriter
		responseData        *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b) // write response using original http.ResponseWriter
	r.responseData.body = b                // capture response body
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode) // write status code using original http.ResponseWriter
	r.responseData.status = statusCode       // capture status code
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		zap.L().Info("Request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		// Create a loggingResponseWriter
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			responseData:   &responseData{},
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(lrw, r)

		var body interface{}
		if lrw.responseData.body != nil {
			json.Unmarshal(lrw.responseData.body, &body)
		}

		// Do stuff here
		zap.L().Debug("Response",
			zap.Int("status", lrw.responseData.status),
			zap.Any("body", body),
		)
	})
}
