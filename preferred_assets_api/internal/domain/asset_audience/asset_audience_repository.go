package assetaudience

type AudienceRepository interface {
	Save(audience Audience) error
	GetByID(id string) (Audience, error)
	GetAll() ([]Audience, error)
	Delete(id string) error
}
