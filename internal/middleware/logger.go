package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{
			"method":      r.Method,
			"path":        r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
		}).Info("Request details")
		t1 := time.Now()
		defer func() {
			logrus.WithFields(logrus.Fields{
				"duration": time.Since(t1).String(),
			}).Info("Request completed")
		}()
		next.ServeHTTP(w, r)
	})
}
