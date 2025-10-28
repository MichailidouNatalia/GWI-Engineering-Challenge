package favourite

import "fmt"

type FavouriteService struct {
	repo FavouriteRepository
}

func NewFavouriteService(r FavouriteRepository) *FavouriteService {
	return &FavouriteService{repo: r}
}

func (s *FavouriteService) AddFavourite(f Favourite) error {
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
