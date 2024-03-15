package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (hr *Handler) checkMiddleware(w http.ResponseWriter, r *http.Request) {
	logrus.Info("in check middleware")
	w.Write([]byte("check middleware for auth users."))
}
