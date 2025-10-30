package domain

import "time"

type Favourite struct {
	UserID    string
	AssetID   string
	CreatedAt time.Time
}
