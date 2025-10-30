package mapper

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// AudienceEntityToDomain converts entity to domain model
func AudienceEntityToDomain(e *entities.AudienceEntity) *domain.Audience {
	return &domain.Audience{
		AssetBase:       AssetBaseEntityToDomain(e.AssetBaseEntity),
		Gender:          e.Gender,
		BirthCountry:    e.BirthCountry,
		AgeGroup:        e.AgeGroup,
		HoursSocial:     e.HoursSocial,
		PurchasesLastMo: e.PurchasesLastMo,
	}
}

// AudienceEntityFromDomain converts domain model to entity
func AudienceEntityFromDomain(a *domain.Audience) entities.AudienceEntity {
	return entities.AudienceEntity{
		AssetBaseEntity: *AssetBaseEntityFromDomain(a.AssetBase),
		Gender:          a.Gender,
		BirthCountry:    a.BirthCountry,
		AgeGroup:        a.AgeGroup,
		HoursSocial:     a.HoursSocial,
		PurchasesLastMo: a.PurchasesLastMo,
	}
}
