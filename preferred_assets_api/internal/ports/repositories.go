package ports

import (
	audience "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset_audience"
	chart "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset_chart"
	insight "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset_insight"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/favourite"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
)

type UserRepository interface {
	GetByID(id string) (*user.User, error)
	Save(user user.User) error
	GetAll() ([]user.User, error)
	Delete(id string) error
	Update(user user.User) error
}

type FavouriteRepository interface {
	Add(f favourite.Favourite) error
	Remove(userID, assetID string) error
	GetByUser(userID string) ([]favourite.Favourite, error)
}

type AudienceRepository interface {
	Save(audience audience.Audience) error
	GetByID(id string) (audience.Audience, error)
	GetAll() ([]audience.Audience, error)
	Delete(id string) error
}

type ChartRepository interface {
	Save(chart chart.Chart) error
	GetByID(id string) (chart.Chart, error)
	GetAll() ([]chart.Chart, error)
	Delete(id string) error
}

type InsightRepository interface {
	Save(insight insight.Insight) error
	GetByID(id string) (insight.Insight, error)
	GetAll() ([]insight.Insight, error)
	Delete(id string) error
}
