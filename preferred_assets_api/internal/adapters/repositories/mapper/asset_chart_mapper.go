package mapper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// ChartEntityToDomain converts entity to domain model
func ChartEntityToDomain(e *entities.ChartEntity) *domain.Chart {
	return &domain.Chart{
		AssetBase:  AssetBaseEntityToDomain(e.AssetBaseEntity),
		AxesTitles: safeUnmarshalStringArray(e.AxesTitles, e.ID, "axes titles"),
		Data:       safeUnmarshalFloatArray(e.Data, e.ID, "chart data"),
	}
}

// ChartEntityFromDomain converts domain model to entity
func ChartEntityFromDomain(c domain.Chart) (*entities.ChartEntity, error) {
	// Serialize slices to JSON
	axesJSON, err := json.Marshal(c.AxesTitles)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal axes titles: %w", err)
	}

	dataJSON, err := json.Marshal(c.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chart data: %w", err)
	}

	return &entities.ChartEntity{
		AssetBaseEntity: *AssetBaseEntityFromDomain(c.AssetBase),
		AxesTitles:      string(axesJSON),
		Data:            string(dataJSON),
	}, nil
}

func safeUnmarshalStringArray(jsonStr, assetID, fieldName string) []string {
	var result []string
	if jsonStr == "" {
		return []string{}
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		log.Printf("Warning: failed to unmarshal %s for asset %s: %v", fieldName, assetID, err)
		return []string{}
	}
	return result
}

func safeUnmarshalFloatArray(jsonStr, assetID, fieldName string) [][]float64 {
	var result [][]float64
	if jsonStr == "" {
		return [][]float64{}
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		log.Printf("Warning: failed to unmarshal %s for asset %s: %v", fieldName, assetID, err)
		return [][]float64{}
	}
	return result
}
