package application

import (
	assetaudience "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset_audience"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type AudienceService struct {
	repo ports.AudienceRepository
}

func NewAudienceService(r ports.AudienceRepository) *AudienceService {
	return &AudienceService{repo: r}
}

func (s *AudienceService) CreateAudience(c assetaudience.Audience) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *AudienceService) GetAudience(id string) (assetaudience.Audience, error) {
	return s.repo.GetByID(id)
}
