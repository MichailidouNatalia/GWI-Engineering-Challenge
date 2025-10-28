package asset

type AssetRepository interface {
	Save(asset Asset) error
	GetByID(id string) (Asset, error)
	GetAll() ([]Asset, error)
	Delete(id string) error
}
