package assetinsight

type InsightRepository interface {
	Save(insight Insight) error
	GetByID(id string) (Insight, error)
	GetAll() ([]Insight, error)
	Delete(id string) error
}
