package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/middleware"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/mapping"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"github.com/go-chi/chi/v5"
)

var _ ports.UserHandler = (*UserHandler)(nil)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(s ports.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// @Summary      Create user
// @Description  Creates a new user account
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateUserRequest true "User info"
// @Success      201 {object} dto.UserResponse
// @Failure      400 {object} map[string]string
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, ok := middleware.GetValidatedBody[dto.CreateUserRequest](r)
	if !ok {
		http.Error(w, "missing validated body", http.StatusBadRequest)
		return
	}

	usr, err := mapping.ToDomain(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.CreateUser(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(created)
}

// Get handles GET /users/{id}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	u, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Fatal(err)
	}
}

// List handles GET /users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Fatal(err)
	}
}

// Get handles Delete /users/{id}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteUser(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	req, ok := middleware.GetValidatedBody[dto.UpdateUserRequest](r)
	if !ok {
		http.Error(w, "missing validated body", http.StatusBadRequest)
		return
	}

	existingUser, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	updatedUser := mapping.UpdateToDomain(existingUser, req)
	if err := h.service.UpdateUser(*updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "updated"}); err != nil {
		http.Error(w, "Failed JSON serialization", http.StatusInternalServerError)
	}
}
