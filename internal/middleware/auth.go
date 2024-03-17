package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/util"
	"github.com/sirupsen/logrus"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := util.ValidateAdminRoleJWT(r)
		if err != nil {
			js, err := json.Marshal(map[string]string{"err": err.Error()})
			if err != nil {
				logrus.Error("Cannot convert error to json")
			}
			http.Error(w, string(js), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func UserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := util.ValidateJWT(r)
		if err != nil {
			js, err := json.Marshal(map[string]string{"err": err.Error()})
			if err != nil {
				logrus.Error("Cannot convert error to json")
			}
			http.Error(w, string(js), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}