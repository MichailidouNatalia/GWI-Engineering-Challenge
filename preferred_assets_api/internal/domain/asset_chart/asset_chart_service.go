package assetchart

type ChartService struct {
	repo ChartRepository
}

func NewChartService(r ChartRepository) *ChartService {
	return &ChartService{repo: r}
}

func (s *ChartService) CreateChart(c Chart) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *ChartService) GetChart(id string) (Chart, error) {
	return s.repo.GetByID(id)
}
