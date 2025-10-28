package application

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

var _ ports.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserService(usrRepo ports.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: usrRepo}
}

func (usrService UserServiceImpl) GetUserByID(id string) (*user.User, error) {
	return usrService.repo.GetByID(id)
}

func (usrService UserServiceImpl) GetAllUsers() ([]user.User, error) {
	return usrService.repo.GetAll()
}

func (usrService UserServiceImpl) CreateUser(usr user.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(hashedPassword)
	usr.CreatedAt = time.Now().UTC()

	return usrService.repo.Save(usr)
}

func (usrService UserServiceImpl) DeleteUser(id string) error {
	return usrService.repo.Delete(id)
}

func (usrService UserServiceImpl) UpdateUser(usr user.User) error {
	usr.UpdateAt = time.Now().UTC()
	return usrService.repo.Update(usr)
}
