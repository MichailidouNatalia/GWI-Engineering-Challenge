package asset

type AssetService struct {
	repo AssetRepository
}

func NewAssetService(r AssetRepository) *AssetService {
	return &AssetService{repo: r}
}

func (s *AssetService) CreateAsset(a Asset) error {
	// You could add domain rules here
	return s.repo.Save(a)
}

func (s *AssetService) GetAsset(id string) (Asset, error) {
	return s.repo.GetByID(id)
}
