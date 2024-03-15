package handler

import (
	"encoding/json"
	"github.com/andyfilya/restapi"
	"net/http"
)

func (hr *Handler) registerNewUser(w http.ResponseWriter, r *http.Request) {
	var newUser restapi.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		newErrWrite(w, http.StatusBadRequest, "bad request")
		return
	}

	userId, err := hr.services.Authorization.NewUser(newUser)
	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendBytes, err := json.Marshal(map[string]interface{}{
		"userId": userId,
	})
	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, "unknown error.")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(sendBytes)
}

func (hr *Handler) signinUser(w http.ResponseWriter, r *http.Request) {
	var usr restapi.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, "error while marshaling username and password.")
		return
	}

	token, err := hr.services.Authorization.NewUserToken(usr.Username, usr.Password)
	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendBytes, err := json.Marshal(token)
	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("token", string(sendBytes))
	w.Write(sendBytes)
}
