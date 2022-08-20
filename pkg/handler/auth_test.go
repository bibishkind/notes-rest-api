package handler

import (
	"bytes"
	"errors"
	"github.com/bibishkind/notes-rest-api/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	// Arrange

	testTable := []struct {
		name                 string
		username             string
		password             string
		id                   int
		err                  error
		requestBody          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Correct data",
			requestBody:          `{"username":"denis", "password":"1234"}`,
			username:             "denis",
			password:             "1234",
			id:                   1,
			err:                  nil,
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Incorrect data",
			requestBody:          `{}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name:                 "Service error",
			requestBody:          `{"username":"denis", "password":"1234"}`,
			username:             "denis",
			password:             "1234",
			id:                   1,
			err:                  errors.New("error message"),
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"error message"}`,
		},
	}

	// Act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			service := new(mocks.MockedService)
			service.On("CreateUser", testCase.username, testCase.password).Return(testCase.id, testCase.err)

			handler := NewHandler(service)
			router := gin.New()
			router.POST("/sign-up", handler.signUp)

			writer := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.requestBody))

			router.ServeHTTP(writer, request)

			// Assert
			assert.Equal(t, writer.Code, testCase.expectedStatusCode)
			assert.Equal(t, writer.Body.String(), testCase.expectedResponseBody)
		})
	}
}

func TestSignIn(t *testing.T) {
	// Arrange

	testTable := []struct {
		name                 string
		username             string
		password             string
		token                string
		err                  error
		requestBody          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Correct data",
			requestBody:          `{"username":"denis", "password":"1234"}`,
			username:             "denis",
			password:             "1234",
			token:                "token",
			err:                  nil,
			expectedStatusCode:   200,
			expectedResponseBody: `{"jwt":"token"}`,
		},
		{
			name:                 "Incorrect data",
			requestBody:          `{}`,
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name:                 "Service error",
			requestBody:          `{"username":"denis", "password":"1234"}`,
			username:             "denis",
			password:             "1234",
			token:                "",
			err:                  errors.New("error message"),
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"error message"}`,
		},
	}

	// Act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			service := new(mocks.MockedService)
			service.On("GenerateJWT", testCase.username, testCase.password).Return(testCase.token, testCase.err)

			handler := NewHandler(service)
			router := gin.New()
			router.POST("/sign-in", handler.signIn)

			writer := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(testCase.requestBody))

			router.ServeHTTP(writer, request)

			// Assert
			assert.Equal(t, writer.Code, testCase.expectedStatusCode)
			assert.Equal(t, writer.Body.String(), testCase.expectedResponseBody)
		})
	}
}
