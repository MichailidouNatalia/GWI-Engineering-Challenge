package application

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type InsightService struct {
	repo ports.InsightRepository
}

func NewInsightService(r ports.InsightRepository) *InsightService {
	return &InsightService{repo: r}
}

func (s *InsightService) CreateInsight(c domain.Insight) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *InsightService) GetInsight(id string) (domain.Insight, error) {
	return s.repo.GetByID(id)
}
