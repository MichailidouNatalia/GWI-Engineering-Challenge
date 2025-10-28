package assetaudience

import "github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/asset"

type Audience struct {
	asset.Asset
	Gender          string
	BirthCountry    string
	AgeGroup        string
	HoursSocial     int
	PurchasesLastMo int
}
