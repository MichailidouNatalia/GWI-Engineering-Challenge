package application

import (
	insight "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/asset_insight"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type InsightService struct {
	repo ports.InsightRepository
}

func NewInsightService(r ports.InsightRepository) *InsightService {
	return &InsightService{repo: r}
}

func (s *InsightService) CreateInsight(c insight.Insight) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *InsightService) GetInsight(id string) (insight.Insight, error) {
	return s.repo.GetByID(id)
}
