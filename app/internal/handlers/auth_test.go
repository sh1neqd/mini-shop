package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testAssignment/internal/config"
	"testAssignment/internal/repositories"
	"testAssignment/internal/services"
	"testAssignment/pkg/client/postgresql"
	"testing"
)

var (
	db, _   = postgresql.NewPostgresDB(config.GetConfig())
	repo    = repositories.NewRepository(db)
	service = services.NewService(repo)
)

func TestSignUp(t *testing.T) {
	requestBody := `{"username": "testuser", "password": "testpassword", "email": "testmail@example.com"}`
	request := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(requestBody))
	responseRecorder := httptest.NewRecorder()

	handler := &Handler{
		services: service,
	}

	handler.signUp(responseRecorder, request, nil)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	expectedResponseBody := `{"message":"user with id=1 successfully created"}`
	if responseRecorder.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body %s, got %s", expectedResponseBody, responseRecorder.Body.String())
	}
}

func TestSignIn(t *testing.T) {
	requestBody := `{"username": "testuser111", "password": "123456"}`
	request := httptest.NewRequest(http.MethodPost, "/signin", strings.NewReader(requestBody))
	responseRecorder := httptest.NewRecorder()

	handler := &Handler{
		services: service,
	}

	handler.signIn(responseRecorder, request, nil)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, responseRecorder.Code)
	}

}
