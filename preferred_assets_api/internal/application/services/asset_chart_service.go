package application

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

type ChartService struct {
	repo ports.ChartRepository
}

func NewChartService(r ports.ChartRepository) *ChartService {
	return &ChartService{repo: r}
}

func (s *ChartService) CreateChart(c domain.Chart) error {
	// Domain rules could go here

	return s.repo.Save(c)
}

func (s *ChartService) GetChart(id string) (domain.Chart, error) {
	return s.repo.GetByID(id)
}
