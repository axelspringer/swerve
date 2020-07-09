package helper

import (
	"net/http"
	"time"

	"github.com/axelspringer/swerve/log"
)

// LoggingMiddleware logs request infos
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Infof("Request starting %s %s %s %s", r.Proto, r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
		end := time.Now()
		diff := end.Sub(start)
		log.Infof("Request finished in %d ms", diff.Nanoseconds()/int64(time.Millisecond))
	})
}
