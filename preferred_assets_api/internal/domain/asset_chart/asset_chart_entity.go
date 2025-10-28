package assetchart

import "github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/asset"

type Chart struct {
	asset.Asset
	AxesTitles []string
	Data       [][]float64
}
