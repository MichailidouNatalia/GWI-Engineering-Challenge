package mapping_test

import (
	"testing"
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/mapping"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// --- FavouriteReqToDomain ---
func TestFavouriteReqToDomain(t *testing.T) {
	// Arrange
	req := dto.FavouriteRequest{UserId: "user1", AssetId: "asset1"}

	// Act
	fav := mapping.FavouriteReqToDomain(req)

	// Assert
	if fav.UserID != "user1" {
		t.Errorf("expected UserID 'user1', got '%s'", fav.UserID)
	}
	if fav.AssetID != "asset1" {
		t.Errorf("expected AssetID 'asset1', got '%s'", fav.AssetID)
	}
}

// --- FavouriteToResponse: nil asset ---
func TestFavouriteToResponse_NilAsset(t *testing.T) {
	// Arrange
	favDomain := domain.Favourite{UserID: "user1", AssetID: "asset1", CreatedAt: time.Now()}

	// Act
	res := mapping.FavouriteToResponse(&favDomain)

	// Assert
	if res.Asset != nil {
		t.Errorf("expected Asset to be nil, got %+v", res.Asset)
	}
	if !res.CreatedAt.Equal(favDomain.CreatedAt) {
		t.Errorf("expected CreatedAt %v, got %v", favDomain.CreatedAt, res.CreatedAt)
	}
}

// --- FavouriteToResponse: Audience ---
func TestFavouriteToResponse_WithAudience(t *testing.T) {
	// Arrange
	audience := &domain.Audience{
		AssetBase:    domain.AssetBase{ID: "a1", Type: domain.AssetTypeAudience, Title: "Audience1"},
		Gender:       "M",
		BirthCountry: "US",
	}
	favDomain := domain.Favourite{UserID: "u1", AssetID: "a1"}
	favDomain.SetAsset(audience)

	// Act
	res := mapping.FavouriteToResponse(&favDomain)

	// Assert
	dtoAsset, ok := res.Asset.(dto.AudienceResponse)
	if !ok {
		t.Fatalf("expected AudienceResponse, got %T", res.Asset)
	}
	if dtoAsset.ID != "a1" {
		t.Errorf("expected ID 'a1', got '%s'", dtoAsset.ID)
	}
	if dtoAsset.Gender != "M" {
		t.Errorf("expected Gender 'M', got '%s'", dtoAsset.Gender)
	}
}

// --- FavouriteToResponse: Chart ---
func TestFavouriteToResponse_WithChart(t *testing.T) {
	// Arrange
	chart := &domain.Chart{
		AssetBase:  domain.AssetBase{ID: "c1", Type: domain.AssetTypeChart, Title: "Chart1"},
		AxesTitles: []string{"X", "Y"},
		Data:       [][]float64{{1, 2}, {3, 4}},
	}
	favDomain := domain.Favourite{UserID: "u1", AssetID: "c1"}
	favDomain.SetAsset(chart)

	// Act
	res := mapping.FavouriteToResponse(&favDomain)

	// Assert
	dtoAsset, ok := res.Asset.(dto.ChartResponse)
	if !ok {
		t.Fatalf("expected ChartResponse, got %T", res.Asset)
	}
	if dtoAsset.ID != "c1" {
		t.Errorf("expected ID 'c1', got '%s'", dtoAsset.ID)
	}
	if len(dtoAsset.Data) != 2 {
		t.Errorf("expected 2 data rows, got %d", len(dtoAsset.Data))
	}
}

// --- FavouriteToResponse: Insight ---
func TestFavouriteToResponse_WithInsight(t *testing.T) {
	// Arrange
	insight := &domain.Insight{
		AssetBase: domain.AssetBase{ID: "i1", Type: domain.AssetTypeInsight, Title: "Insight1"},
		Text:      "Some insight text",
	}
	favDomain := domain.Favourite{UserID: "u1", AssetID: "i1"}
	favDomain.SetAsset(insight)

	// Act
	res := mapping.FavouriteToResponse(&favDomain)

	// Assert
	dtoAsset, ok := res.Asset.(dto.InsightResponse)
	if !ok {
		t.Fatalf("expected InsightResponse, got %T", res.Asset)
	}
	if dtoAsset.Text != "Some insight text" {
		t.Errorf("expected Text 'Some insight text', got '%s'", dtoAsset.Text)
	}
}

// --- FavouritesToResponse: mixed assets and nil ---
func TestFavouritesToResponse_MixedAssets(t *testing.T) {
	// Arrange
	aud := &domain.Audience{AssetBase: domain.AssetBase{ID: "a1", Type: domain.AssetTypeAudience, Title: "Audience1"}}
	chart := &domain.Chart{AssetBase: domain.AssetBase{ID: "c1", Type: domain.AssetTypeChart, Title: "Chart1"}}
	fav1 := domain.Favourite{UserID: "u1", AssetID: "a1"}
	fav1.SetAsset(aud)
	fav2 := domain.Favourite{UserID: "u2", AssetID: "c1"}
	fav2.SetAsset(chart)
	fav3 := domain.Favourite{UserID: "u3", AssetID: "i1"} // nil asset
	favs := []domain.Favourite{fav1, fav2, fav3}

	// Act
	responses := mapping.FavouritesToResponse(favs)

	// Assert
	if len(responses) != 3 {
		t.Fatalf("expected 3 responses, got %d", len(responses))
	}
	if responses[0].Asset == nil || responses[1].Asset == nil {
		t.Error("expected first two assets to be non-nil")
	}
	if responses[2].Asset != nil {
		t.Error("expected third asset to be nil")
	}
}

// --- Unhappy path: nil Favourite pointer (domain object) ---
func TestFavouriteToResponse_NilDomainObject(t *testing.T) {
	// Arrange
	var fav *domain.Favourite // nil pointer

	// Act
	res := mapping.FavouriteToResponse(fav)

	// Assert
	if res.UserID != "" || res.AssetID != "" || res.Asset != nil {
		t.Errorf("expected empty FavouriteResponse, got %+v", res)
	}
}
