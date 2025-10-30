package mapper

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// AssetBaseDomainToEntity domain to entity mapping
func AssetBaseEntityFromDomain(a domain.AssetBase) *entities.AssetBaseEntity {
	return &entities.AssetBaseEntity{
		ID:          a.ID,
		Type:        int(a.Type),
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
