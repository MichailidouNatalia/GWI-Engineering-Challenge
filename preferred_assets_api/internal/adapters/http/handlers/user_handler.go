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

// Create creates a new user
// swagger:operation POST /users users createUser
//
// Create User
// ---
// responses:
//
//	201: NoContentResponse
//	400: ValidationErrorResponse
//	405: MethodErrorResponse
//	500: ServerErrorResponse
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

	usr, err := mapping.UserReqToDomain(req)
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
}

// Get retrieves a user by ID
// swagger:operation GET /users/{id} users getUser
//
// Get User
// ---
// responses:
//
//	200: UserResponse
//	400: ValidationErrorResponse
//	404: NotFoundResponse
//	405: MethodErrorResponse
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

	usr := mapping.DomainToUserRes(*u)
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		log.Fatal(err)
	}
}

// List retrieves all users
// swagger:operation GET /users users listUsers
//
// List Users
// ---
// responses:
//
//	200: UserListResponse
//	405: MethodErrorResponse
//	500: ServerErrorResponse
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

	usersResponse := mapping.UserReqToResponseList(users)
	if err := json.NewEncoder(w).Encode(usersResponse); err != nil {
		log.Fatal(err)
	}
}

// Delete removes a user by ID
// swagger:operation DELETE /users/{id} users deleteUser
//
// Delete User
// ---
// responses:
//
//	200: NoContentResponse
//	400: ValidationErrorResponse
//	404: NotFoundResponse
//	405: MethodErrorResponse
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

// Update modifies a user by ID
// swagger:operation PUT /users/{id} users updateUser
//
// Update User
// ---
// responses:
//
//	200: UpdateSuccessResponse
//	400: ValidationErrorResponse
//	404: NotFoundResponse
//	405: MethodErrorResponse
//	500: ServerErrorResponse
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

	updatedUser := mapping.UpdateReqToDomain(existingUser, req)
	if err := h.service.UpdateUser(*updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "updated"}); err != nil {
		http.Error(w, "Failed JSON serialization", http.StatusInternalServerError)
	}
}

// GetFavourites retrieves user favourites
// swagger:operation GET /users/{id}/favourites users getUserFavourites
//
// Get User Favourites
// ---
// responses:
//
//	200: FavouriteListResponse
//	400: ValidationErrorResponse
//	404: NotFoundResponse
//	405: MethodErrorResponse
//	500: ServerErrorResponse
func (h *UserHandler) GetFavourites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	favourites, err := h.service.GetFavouritesByUser(id)
	if err != nil {
		http.Error(w, "favourites not found", http.StatusNotFound)
		return
	}

	response := mapping.FavouritesToResponse(favourites)

	jsonBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("JSON marshaling error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Favourites JSON response: %s", string(jsonBytes))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
