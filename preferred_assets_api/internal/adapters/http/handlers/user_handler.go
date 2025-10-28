package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/user"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Create handles POST /users
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req user.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(created)
}

// Get handles GET /users/{id}?id=123
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
