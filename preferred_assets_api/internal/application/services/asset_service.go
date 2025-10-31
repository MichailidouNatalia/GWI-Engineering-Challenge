package services

import (
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/mapper"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

var _ ports.AssetService = (*AssetServiceImpl)(nil)

type AssetServiceImpl struct {
	assetRepo ports.AssetRepository
}

func NewAssetService(assetRepo ports.AssetRepository) *AssetServiceImpl {
	return &AssetServiceImpl{
		assetRepo: assetRepo}
}

// CreateAsset implements ports.AssetService.
func (assetService *AssetServiceImpl) CreateAsset(asset domain.Asset) error {
	asset.SetCreatedAt(time.Now().UTC())
	assetEntity, err := mapper.AssetEntityFromDomain(asset)
	if err != nil {
		return err
	}

	return assetService.assetRepo.Save(assetEntity)

}

// DeleteAsset implements ports.AssetService.
func (assetService *AssetServiceImpl) DeleteAsset(id string) error {
	return assetService.assetRepo.Delete(id)
}
