package assetaudience

type AudienceService struct {
	repo AudienceRepository
}

func NewAudienceService(r AudienceRepository) *AudienceService {
	return &AudienceService{repo: r}
}

func (s *AudienceService) CreateAudience(c Audience) error {
	// Domain rules could go here
	return s.repo.Save(c)
}

func (s *AudienceService) GetAudience(id string) (Audience, error) {
	return s.repo.GetByID(id)
}
