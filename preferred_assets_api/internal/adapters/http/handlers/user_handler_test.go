package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/handlers"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
)

// ------------------------
// Mocks
// ------------------------
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(u domain.User) error { return m.Called(u).Error(0) }
func (m *MockUserService) GetUserByID(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserService) GetAllUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
func (m *MockUserService) UpdateUser(u domain.User) error { return m.Called(u).Error(0) }
func (m *MockUserService) DeleteUser(id string) error     { return m.Called(id).Error(0) }

// Middleware mock
type MockBodyGetter struct{ MockedBody any }

func (m MockBodyGetter) GetValidatedBody(r *http.Request) (any, bool) { return m.MockedBody, true }

// Helper for requests
func performRequest(handlerFunc http.HandlerFunc, method, url string, body any, ctxParams map[string]string) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		json.NewEncoder(&buf).Encode(body)
	}
	req := httptest.NewRequest(method, url, &buf)
	rctx := chi.NewRouteContext()
	for k, v := range ctxParams {
		rctx.URLParams.Add(k, v)
	}
	req = req.WithContext(rctx.Context())
	w := httptest.NewRecorder()
	handlerFunc(w, req)
	return w
}

// ------------------------
// Tests
// ------------------------
func TestUserHandler(t *testing.T) {
	mockService := new(MockUserService)
	handler := handlers.NewUserHandler(mockService)

	// ------------------------
	// Create User
	// ------------------------
	t.Run("Create_Happy", func(t *testing.T) {
		handlers.Body = MockBodyGetter{MockedBody: dto.CreateUserRequest{Name: "Alice", Email: "a@b.com", Password: "pwd12345"}}
		mockService.On("CreateUser", mock.AnythingOfType("domain.User")).Return(nil)

		w := performRequest(handler.Create, http.MethodPost, "/users", handlers.Body.GetValidatedBody(nil), nil)
		if w.Code != http.StatusCreated {
			t.Errorf("expected 201, got %d", w.Code)
		}
	})

	t.Run("Create_MiddlewareFails", func(t *testing.T) {
		handlers.Body = MockBodyGetter{MockedBody: nil}
		w := performRequest(handler.Create, http.MethodPost, "/users", nil, nil)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected 400, got %d", w.Code)
		}
	})

	// ------------------------
	// Get User
	// ------------------------
	t.Run("Get_Happy", func(t *testing.T) {
		mockService.On("GetUserByID", "123").Return(&domain.User{Id: "123", Name: "Alice"}, nil)
		w := performRequest(handler.Get, http.MethodGet, "/users/123", nil, map[string]string{"id": "123"})
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Get_NotFound", func(t *testing.T) {
		mockService.On("GetUserByID", "999").Return(nil, errors.New("not found"))
		w := performRequest(handler.Get, http.MethodGet, "/users/999", nil, map[string]string{"id": "999"})
		if w.Code != http.StatusNotFound {
			t.Errorf("expected 404, got %d", w.Code)
		}
	})

	// ------------------------
	// List Users
	// ------------------------
	t.Run("List_Happy", func(t *testing.T) {
		mockService.On("GetAllUsers").Return([]domain.User{{Id: "1", Name: "Alice"}}, nil)
		w := performRequest(handler.List, http.MethodGet, "/users", nil, nil)
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("List_ServiceFails", func(t *testing.T) {
		mockService.On("GetAllUsers").Return(nil, errors.New("service error"))
		w := performRequest(handler.List, http.MethodGet, "/users", nil, nil)
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})

	// ------------------------
	// Update User
	// ------------------------
	t.Run("Update_Happy", func(t *testing.T) {
		reqDTO := dto.UpdateUserRequest{Name: "Alice Updated"}
		handlers.Body = MockBodyGetter{MockedBody: reqDTO}
		mockService.On("UpdateUser", mock.AnythingOfType("domain.User")).Return(nil)

		w := performRequest(handler.Update, http.MethodPut, "/users/123", reqDTO, map[string]string{"id": "123"})
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Update_ServiceFails", func(t *testing.T) {
		reqDTO := dto.UpdateUserRequest{Name: "Alice Updated"}
		handlers.Body = MockBodyGetter{MockedBody: reqDTO}
		mockService.On("UpdateUser", mock.AnythingOfType("domain.User")).Return(errors.New("service error"))

		w := performRequest(handler.Update, http.MethodPut, "/users/123", reqDTO, map[string]string{"id": "123"})
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})

	// ------------------------
	// Delete User
	// ------------------------
	t.Run("Delete_Happy", func(t *testing.T) {
		mockService.On("DeleteUser", "123").Return(nil)
		w := performRequest(handler.Delete, http.MethodDelete, "/users/123", nil, map[string]string{"id": "123"})
		if w.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", w.Code)
		}
	})

	t.Run("Delete_ServiceFails", func(t *testing.T) {
		mockService.On("DeleteUser", "123").Return(errors.New("service error"))
		w := performRequest(handler.Delete, http.MethodDelete, "/users/123", nil, map[string]string{"id": "123"})
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected 500, got %d", w.Code)
		}
	})
}
