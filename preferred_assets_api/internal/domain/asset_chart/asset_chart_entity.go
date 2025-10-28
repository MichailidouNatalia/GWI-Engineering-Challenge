package assetchart

import "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset"

type Chart struct {
	asset.Asset
	AxesTitles []string
	Data       [][]float64
}
