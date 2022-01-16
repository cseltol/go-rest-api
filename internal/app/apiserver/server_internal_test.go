package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cseltol/go-rest-api/internal/app/model"
	"github.com/cseltol/go-rest-api/internal/app/store/teststore"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": "user@example.com",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
				// and also missing password
			},
			expectedCode: http.StatusUnprocessableEntity,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("secret")))
	testCases := []struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email": u.Email,
				"password": u.Passowrd,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"email": "invalid",
				"password": u.Passowrd,
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email": u.Email,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},

	} 

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}