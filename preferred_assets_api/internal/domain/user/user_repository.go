package user

type UserRepository interface {
	GetByID(id string) (User, error)
	Save(user User) error
	GetAll() ([]User, error)
	Delete(id string) error
}
