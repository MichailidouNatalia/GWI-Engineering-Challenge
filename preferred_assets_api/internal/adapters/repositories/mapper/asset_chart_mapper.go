package mapper

import (
	"encoding/json"
	"fmt"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// ChartEntityToDomain converts entity to domain model
func ChartEntityToDomain(e entities.ChartEntity) (*domain.Chart, error) {
	var axesTitles []string
	var data [][]float64

	// Deserialize axes titles
	if err := json.Unmarshal([]byte(e.AxesTitles), &axesTitles); err != nil {
		return nil, fmt.Errorf("failed to unmarshal axes titles: %w", err)
	}

	// Deserialize chart data
	if err := json.Unmarshal([]byte(e.Data), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal chart data: %w", err)
	}

	return &domain.Chart{
		AssetBase:  AssetBaseEntityToDomain(e.AssetBaseEntity),
		AxesTitles: axesTitles,
		Data:       data,
	}, nil
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
