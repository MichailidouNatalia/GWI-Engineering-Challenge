package application

import (
	"fmt"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/favourite"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type FavouriteService struct {
	repo ports.FavouriteRepository
}

func NewFavouriteService(r ports.FavouriteRepository) *FavouriteService {
	return &FavouriteService{repo: r}
}

func (s *FavouriteService) AddFavourite(f favourite.Favourite) error {
	// Domain logic: prevent duplicates, enforce limits, etc.
	existing, _ := s.repo.GetByUser(f.UserID)
	for _, e := range existing {
		if e.AssetID == f.AssetID {
			return fmt.Errorf("asset already favourited")
		}
	}
	return s.repo.Add(f)
}

func (s *FavouriteService) RemoveFavourite(userID, assetID string) error {
	return s.repo.Remove(userID, assetID)
}
