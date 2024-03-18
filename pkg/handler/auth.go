package handler

import (
	"encoding/json"
	"github.com/andyfilya/restapi"
	"net/http"
)

type toSendRegister struct {
	UserId int
}

// @Summary RegisterNewUser
// @Tags auth
// @Description create new user
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body restapi.User true "account info"
// @Success 200 {object} toSendRegister
// @Failure 400,404 {object} errorMessage
// @Failure 500 {object} errorMessage
// @Failure default {object} errorMessage
// @Router /api/auth/register [post]
func (hr *Handler) registerNewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
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

	sendBytes, err := json.Marshal(toSendRegister{
		UserId: userId,
	})

	if err != nil {
		newErrWrite(w, http.StatusInternalServerError, "unknown error")
		return
	}

	w.Header().Set("Content-Type", "applicatopn/json")
	w.Write(sendBytes)
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/auth/signin [post]
func (hr *Handler) signinUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		newErrWrite(w, http.StatusBadRequest, "bad method")
		return
	}
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

	w.Header().Set("Content-Type", "applicatopn/json")
	w.Header().Set("token", string(sendBytes))
	w.Write(sendBytes)
}
