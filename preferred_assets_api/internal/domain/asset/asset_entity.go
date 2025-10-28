package asset

import "time"

type Asset struct {
	ID          string
	Type        AssetType
	Title       string
	Description string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
