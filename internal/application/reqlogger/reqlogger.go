//nolint:all
package reqlogger

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int
	body   *bytes.Buffer
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		status:         http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	rw.body.Write(b)
	return size, err
}

func LoggerMDW(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := logger.NewLogger(db)
			res := newResponseWriter(w)

			next.ServeHTTP(res, r)

			reqMsg := fmt.Sprintf("Verb: %s | URL: %s | Bytes: %d", r.Method, r.URL.Path, r.ContentLength)
			logger.Info(reqMsg)

			resMsg := fmt.Sprintf("Status: %d | Bytes: %d | Body: %s", res.status, res.size, res.body.String())
			if res.status >= http.StatusBadRequest {
				logger.Error(resMsg)
			} else {
				logger.Info(resMsg)
			}
		})
	}
}
