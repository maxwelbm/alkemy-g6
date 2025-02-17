//nolint:all
package reqlogger

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/maxwelbm/alkemy-g6/pkg/logger"
)

func LoggerMDW(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := logger.NewLogger(db)
			next.ServeHTTP(w, r)
			msg := fmt.Sprintf("Verb: %s | URL: %s | Bytes: %d", r.Method, r.URL.Path, r.ContentLength)
			logger.Info(msg)
		})
	}
}
