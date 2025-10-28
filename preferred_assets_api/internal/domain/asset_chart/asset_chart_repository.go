package assetchart

type ChartRepository interface {
	Save(chart Chart) error
	GetByID(id string) (Chart, error)
	GetAll() ([]Chart, error)
	Delete(id string) error
}
