package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	auth = "Authorization"
)

func (hr *Handler) middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get(auth)
		hr.logger.Infof("user token : [%v]", authToken)
		if authToken == "" {
			newErrWrite(w, http.StatusUnauthorized, "you need to login in this endpoint | bad token.")
			return
		}
		authTokenSplit := strings.Split(authToken, " ")
		if len(authTokenSplit) != 2 || authTokenSplit[0] != "Bearer" || len(authTokenSplit[1]) == 0 {
			newErrWrite(w, http.StatusUnauthorized, "you need to login in this endpoint | bad token.")
			return
		}

		userId, err := hr.services.ParseUserToken(authTokenSplit[1])
		if err != nil {
			newErrWrite(w, http.StatusUnauthorized, "bad token.")
			return
		}
		hr.logger.Infof("user id in middleware : [%d]", userId)
		ctx := context.WithValue(r.Context(), "userId", userId)
		r.WithContext(ctx)

		next(w, r)
	}
}

func (hr *Handler) adminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get(auth)
		if authToken == "" {
			newErrWrite(w, http.StatusUnauthorized, "you need to login in this endpoint | bad token.")
			return
		}
		authTokenSplit := strings.Split(authToken, " ")
		if len(authTokenSplit) != 2 || authTokenSplit[0] != "Bearer" || len(authTokenSplit[1]) == 0 {
			newErrWrite(w, http.StatusUnauthorized, "you need to login in this endpoint | bad token.")
			return
		}

		userRole, err := hr.services.ParseAdminToken(authTokenSplit[1])
		if err != nil {
			newErrWrite(w, http.StatusUnauthorized, "bad token")
			return
		}
		hr.logger.Infof("user role in middleware : [%s]", userRole)
		ctx := context.WithValue(r.Context(), "userId", userRole)
		r.WithContext(ctx)

		next(w, r)
	}
}

func (hr *Handler) setIdMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "req-id", id)))
	}
}

func (hr *Handler) loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := hr.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value("req-id"),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		now := time.Now()

		next(w, r)

		logger.Infof("completed with %d", time.Since(now))
	}
}
