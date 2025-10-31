package mapper

import (
	"fmt"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

func AssetEntityToDomain(entity entities.AssetEntity) (domain.Asset, error) {
	if entity == nil {
		return nil, nil
	}

	switch e := entity.(type) {
	case *entities.AudienceEntity:
		if e == nil {
			return nil, nil
		}
		return AudienceEntityToDomain(e), nil
	case *entities.ChartEntity:
		if e == nil {
			return nil, nil
		}
		return ChartEntityToDomain(e), nil
	case *entities.InsightEntity:
		if e == nil {
			return nil, nil
		}
		return InsightEntityToDomain(e), nil
	default:
		return nil, fmt.Errorf("unknown asset type: %v", entity.GetType())
	}
}

func AssetEntityFromDomain(asset domain.Asset) (entities.AssetEntity, error) {
	if asset == nil {
		return nil, nil
	}

	switch a := asset.(type) {
	case *domain.Audience:
		if a == nil {
			return nil, nil
		}
		return AudienceEntityFromDomain(a), nil
	case *domain.Chart:
		if a == nil {
			return nil, nil
		}
		return ChartEntityFromDomain(a), nil
	case *domain.Insight:
		if a == nil {
			return nil, nil
		}
		return InsightEntityFromDomain(a), nil
	default:
		return nil, fmt.Errorf("unknown asset type: %v", asset.GetType())
	}
}

// AssetBaseDomainToEntity domain to entity mapping
func AssetBaseEntityFromDomain(a domain.AssetBase) *entities.AssetBaseEntity {
	return &entities.AssetBaseEntity{
		ID:          a.ID,
		Type:        entities.AssetType(a.Type),
		Title:       a.Title,
		Description: a.Description,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

// AssetBaseEntityToDomain entity to domain mapping
func AssetBaseEntityToDomain(e entities.AssetBaseEntity) domain.AssetBase {
	return domain.AssetBase{
		ID:          e.ID,
		Type:        domain.AssetType(e.Type),
		Title:       e.Title,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
