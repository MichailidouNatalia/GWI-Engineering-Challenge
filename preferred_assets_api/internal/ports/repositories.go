package ports

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
)

type UserRepository interface {
	GetByID(id string) (*entities.UserEntity, error)
	Save(user entities.UserEntity) error
	GetAll() ([]entities.UserEntity, error)
	Delete(id string) error
	Update(user entities.UserEntity) error
	GetFavouritesByID(id string) ([]entities.FavouriteEntity, error)
}

type FavouriteRepository interface {
	Add(f entities.FavouriteEntity) error
	Delete(userID, assetID string) error
	GetByUserID(userID string) ([]string, error)
	Exists(userID, assetID string) (bool, error)
}

type AudienceRepository interface {
	Save(audience domain.Audience) error
	GetByID(id string) (domain.Audience, error)
	GetAll() ([]domain.Audience, error)
	Delete(id string) error
	Update(audience domain.Audience) error
}

type ChartRepository interface {
	Save(chart domain.Chart) error
	GetByID(id string) (domain.Chart, error)
	GetAll() ([]domain.Chart, error)
	Delete(id string) error
	Update(chart domain.Chart) error
}

type InsightRepository interface {
	Save(insight domain.Insight) error
	GetByID(id string) (domain.Insight, error)
	GetAll() ([]domain.Insight, error)
	Delete(id string) error
	Update(insight domain.Insight)
}
