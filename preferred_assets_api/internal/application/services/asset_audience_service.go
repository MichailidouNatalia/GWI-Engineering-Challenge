package application

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type AudienceService struct {
	repo ports.AudienceRepository
}

func NewAudienceService(r ports.AudienceRepository) *AudienceService {
	return &AudienceService{repo: r}
}

func (s *AudienceService) CreateAudience(c domain.Audience) error {

	return s.repo.Save(c)
}

func (s *AudienceService) GetAudience(id string) (domain.Audience, error) {
	return s.repo.GetByID(id)
}
