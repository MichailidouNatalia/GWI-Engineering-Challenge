package handlers

import (
	"net/http"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/middleware"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/mapping"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"github.com/go-chi/chi/v5"
)

var _ ports.FavouriteHandler = (*FavouriteHandler)(nil)

type FavouriteHandler struct {
	service ports.FavouriteService
}

func NewFavouriteHandler(s ports.FavouriteService) *FavouriteHandler {
	return &FavouriteHandler{service: s}
}

// Create adds a favourite for a user
// swagger:operation POST /favourites favourites createFavourite
//
// Create Favourite
// ---
// responses:
//
//	201: NoContentResponse
//	400: ValidationErrorResponse
//	405: MethodErrorResponse
//	500: ServerErrorResponse
func (f *FavouriteHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, ok := middleware.GetValidatedBody[dto.FavouriteRequest](r)
	if !ok {
		http.Error(w, "missing validated body", http.StatusBadRequest)
		return
	}

	favourite := mapping.FavouriteReqToDomain(req)

	err := f.service.CreateFavourite(favourite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Delete removes a favourite for a user
// swagger:operation DELETE /favourites/{userId}/{assetId} favourites deleteFavourite
//
// Delete Favourite
// ---
// responses:
//
//	200: NoContentResponse
//	400: ValidationErrorResponse
//	404: NotFoundResponse
//	405: MethodErrorResponse
func (f *FavouriteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	usrId := chi.URLParam(r, "userId")
	if usrId == "" {
		http.Error(w, "missing user id", http.StatusBadRequest)
		return
	}

	assetId := chi.URLParam(r, "assetId")
	if assetId == "" {
		http.Error(w, "missing asset id", http.StatusBadRequest)
		return
	}

	err := f.service.DeleteFavourite(usrId, assetId)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
}
