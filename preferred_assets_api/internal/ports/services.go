package ports

import (
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
)

type UserService interface {
	CreateUser(user user.User) error
	GetUserByID(id string) (*user.User, error)
	GetAllUsers() ([]user.User, error)
	UpdateUser(user user.User) error
	DeleteUser(id string) error
}
