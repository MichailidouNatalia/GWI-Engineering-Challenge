package dto

type FavouriteRequest struct {
	UserId  string `json:"userId" validate:"required"`
	AssetId string `json:"assetId" validate:"required"`
}
