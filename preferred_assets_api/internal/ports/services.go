package ports

import "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"

type UserService interface {
	CreateUser(user domain.User) error
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	UpdateUser(user domain.User) error
	DeleteUser(id string) error
	GetFavouritesByUser(id string) ([]domain.Favourite, error)
}

type AudienceService interface {
	CreateAudience(audience domain.Audience) error
	GetAudienceByID(id string) (*domain.Audience, error)
	GetAllAudiences() ([]domain.Audience, error)
	UpdateAudience(audience domain.Audience) error
	DeleteAudience(id string) error
}

type ChartService interface {
	CreateChart(chart domain.Chart) error
	GetChartByID(id string) (*domain.Chart, error)
	GetAllCharts() ([]domain.Chart, error)
	UpdateChart(chart domain.Chart) error
	DeleteChart(id string) error
}

type InsightService interface {
	CreateInsight(insight domain.Insight) error
	GetInsightByID(id string) (*domain.Insight, error)
	GetAllInsights() ([]domain.Insight, error)
	UpdateInsight(insight domain.Insight) error
	DeleteInsight(id string) error
}

type FavouriteService interface {
	CreateFavourite(favourite domain.Favourite) error
	DeleteFavourite(userId string, assetId string) error
}
