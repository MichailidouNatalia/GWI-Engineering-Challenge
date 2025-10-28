package assetinsight

type InsightService struct {
	repo InsightRepository
}

func NewInsightService(r InsightRepository) *InsightService {
	return &InsightService{repo: r}
}

func (s *InsightService) CreateInsight(c Insight) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *InsightService) GetInsight(id string) (Insight, error) {
	return s.repo.GetByID(id)
}
