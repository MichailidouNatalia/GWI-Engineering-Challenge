package mapper

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

// ToDomain converts entity to domain model
func FavouriteEntityToDomain(e entities.FavouriteEntity) domain.Favourite {
	return domain.Favourite{
		UserID:    e.UserId,
		AssetID:   e.AssetId,
		CreatedAt: e.CreatedAt,
	}
}

// FromDomain converts domain model to entity
func FavouriteEntityFromDomain(favourite domain.Favourite) entities.FavouriteEntity {
	return entities.FavouriteEntity{
		UserId:    favourite.UserID,
		AssetId:   favourite.AssetID,
		CreatedAt: favourite.CreatedAt,
	}
}

func FavouriteEntityToDomainList(favourites []entities.FavouriteEntity) []domain.Favourite {
	ent := make([]domain.Favourite, len(favourites))
	for i, favourite := range favourites {
		ent[i] = FavouriteEntityToDomain(favourite)
	}
	return ent
}
