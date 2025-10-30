package mapping

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

func FavouriteReqToDomain(req dto.FavouriteRequest) domain.Favourite {
	return domain.Favourite{
		UserID:  req.UserId,
		AssetID: req.AssetId,
	}
}
