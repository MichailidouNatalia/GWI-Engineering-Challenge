// internal/transport/http/mapping/asset_mapping.go
package mapping

import (
	"fmt"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

func AssetReqToDomain(req dto.AssetRequest) (domain.Asset, error) {
	assetType, err := domain.ParseAssetType(req.Type)
	if err != nil {
		return nil, err
	}

	base := domain.AssetBase{
		ID:          req.ID,
		Type:        assetType,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}

	switch assetType {
	case domain.AssetTypeAudience:
		return &domain.Audience{
			AssetBase:       base,
			Gender:          safeDerefString(req.Gender),
			BirthCountry:    safeDerefString(req.BirthCountry),
			AgeGroup:        safeDerefString(req.AgeGroup),
			HoursSocial:     safeDerefInt(req.HoursSocial),
			PurchasesLastMo: safeDerefInt(req.PurchasesLastMo),
		}, nil

	case domain.AssetTypeChart:
		return &domain.Chart{
			AssetBase:  base,
			AxesTitles: req.AxesTitles,
			Data:       req.Data,
		}, nil

	case domain.AssetTypeInsight:
		return &domain.Insight{
			AssetBase: base,
			// Add insight-specific fields as needed
		}, nil

	default:
		return nil, fmt.Errorf("unsupported asset type: %s", req.Type)
	}
}

// Helper functions for safe dereferencing
func safeDerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func safeDerefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
