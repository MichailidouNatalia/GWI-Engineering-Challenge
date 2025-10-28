package user

type UserService struct {
	repo UserRepository
}

func (s UserService) GetUserByID(id string) (any, any) {
	return s.repo.GetByID(id)
}

func (s UserService) GetAllUsers() (any, any) {
	return s.repo.GetAll()
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(c User) error {
	return s.repo.Save(c)
}
