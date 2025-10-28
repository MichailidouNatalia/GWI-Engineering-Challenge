package assetinsight

import "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset"

type Insight struct {
	asset.Asset
	Text string
}
