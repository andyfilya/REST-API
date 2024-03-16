package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	auth = "Authorization"
)

func (hr *Handler) middlewareAuth(next http.HandlerFunc) http.HandlerFunc {
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
