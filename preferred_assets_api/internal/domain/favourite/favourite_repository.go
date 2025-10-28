package favourite

type FavouriteRepository interface {
	Add(f Favourite) error
	Remove(userID, assetID string) error
	GetByUser(userID string) ([]Favourite, error)
}
