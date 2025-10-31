package handlers

import (
	"net/http"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/middleware"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/mapping"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"github.com/go-chi/chi/v5"
)

var _ ports.AssetHandler = (*AssetHandler)(nil)

type AssetHandler struct {
	service ports.AssetService
}

func NewAssetHandler(s ports.AssetService) *AssetHandler {
	return &AssetHandler{service: s}
}

// Create implements ports.AssetHandler.
func (h *AssetHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req, ok := middleware.GetValidatedBody[dto.AssetRequest](r)
	if !ok {
		http.Error(w, "missing validated body", http.StatusBadRequest)
		return
	}

	asset, err := mapping.AssetReqToDomain(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateAsset(asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// Optionally return the created asset
	// json.NewEncoder(w).Encode(created)
}

// Delete implements ports.AssetHandler.
func (h *AssetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	assetID := chi.URLParam(r, "assetId")
	if assetID == "" {
		http.Error(w, "missing asset id", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteAsset(assetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// List implements ports.AssetHandler.
func (h *AssetHandler) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Get implements ports.AssetHandler (if you have a Get method in your interface)
func (h *AssetHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	assetID := chi.URLParam(r, "assetId")
	if assetID == "" {
		http.Error(w, "missing asset id", http.StatusBadRequest)
		return
	}

	// This would require a GetAsset method in your service interface
	// asset, err := h.service.GetAsset(assetID)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusNotFound)
	//     return
	// }

	// json.NewEncoder(w).Encode(asset)
}
