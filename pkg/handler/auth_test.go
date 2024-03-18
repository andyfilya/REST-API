package handler

import (
	"bytes"
	"errors"
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/pkg/service"
	service_mocks "github.com/andyfilya/restapi/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler_newUserRegister(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, user restapi.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            restapi.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "usernameeee", "user_role": "user", "password": "fkljfajFKJnfjknDJKjfnjk"}`,
			inputUser: restapi.User{
				Username: "usernameeee",
				Password: "fkljfajFKJnfjknDJKjfnjk",
				Role:     "user",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user restapi.User) {
				r.EXPECT().NewUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"UserId":1}`,
		},
		{
			name:      "Valid Password",
			inputBody: `{"username": "username", "user_role": "user", "password": "123"}`,
			inputUser: restapi.User{
				Username: "username",
				Password: "123",
				Role:     "user",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user restapi.User) {
				r.EXPECT().NewUser(user).Return(-1, errors.New("not valid password"))
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"not valid password"}`,
		},
		{
			name:      "Not user",
			inputBody: `{"username": "username", "user_role": "role", "password": "kfkfkjjJFJjfkKFKkfkjFJJFkjfjJFN"}`,
			inputUser: restapi.User{
				Username: "username",
				Password: "kfkfkjjJFJjfkKFKkfkjFJJFkjfjJFN",
				Role:     "role",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, user restapi.User) {
				r.EXPECT().NewUser(user).Return(-1, errors.New("you are not admin"))
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"you are not admin"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			logger := &logrus.Logger{
				Out:   os.Stderr,
				Level: logrus.DebugLevel,
				Formatter: &logrus.JSONFormatter{
					TimestampFormat: "2006-01-02 15:04:05",
					PrettyPrint:     true,
				},
			}

			services := &service.Service{Authorization: repo}
			handler := Handler{services, logger}

			// Init Endpoint
			m := http.NewServeMux()
			m.HandleFunc("/auth/register", handler.registerNewUser)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/register",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			m.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
