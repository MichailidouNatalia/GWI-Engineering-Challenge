package assetinsight

import "github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/asset"

type Insight struct {
	asset.Asset
	Text string
}
