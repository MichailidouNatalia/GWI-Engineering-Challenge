package application

import (
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/mapper"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"golang.org/x/crypto/bcrypt"
)

var _ ports.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(usrRepo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: usrRepo}
}

func (usrService UserServiceImpl) GetUserByID(id string) (*domain.User, error) {
	user, err := usrService.repo.GetByID(id)
	return mapper.UserEntityToDomain(user), err
}

func (usrService UserServiceImpl) GetAllUsers() ([]domain.User, error) {
	users, err := usrService.repo.GetAll()
	userList := mapper.UserEntintyToDomainList(users)
	return userList, err
}

func (usrService UserServiceImpl) CreateUser(usr domain.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(hashedPassword)
	usr.CreatedAt = time.Now().UTC()
	user := mapper.UserEntityFromDomain(usr)

	return usrService.repo.Save(user)
}

func (usrService UserServiceImpl) DeleteUser(id string) error {
	return usrService.repo.Delete(id)
}

func (usrService UserServiceImpl) UpdateUser(usr domain.User) error {
	usr.UpdatedAt = time.Now().UTC()
	user := mapper.UserEntityFromDomain(usr)
	return usrService.repo.Update(user)
}

func (usrService UserServiceImpl) GetFavouritesByUser(id string) ([]domain.Favourite, error) {
	favs, err := usrService.repo.GetFavouritesByID(id)
	favsList := mapper.FavouriteEntityToDomainList(favs)
	return favsList, err
}
