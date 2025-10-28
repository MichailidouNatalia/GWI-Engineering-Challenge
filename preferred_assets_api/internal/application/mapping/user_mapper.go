package mapping

import (
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
	"github.com/google/uuid"
)

func ToDomain(req dto.CreateUserRequest) (user.User, error) {

	return user.User{
		ID:        uuid.NewString(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}, nil
}

func UpdateToDomain(existingUser *user.User, req dto.UpdateUserRequest) *user.User {
	// Apply updates to existing domain entity
	if req.Email != "" {
		existingUser.Email = req.Email
	}
	if req.Name != "" {
		existingUser.Name = req.Name
	}
	if req.Password != "" {
		existingUser.Password = req.Password
	}
	return existingUser
}

func ToResponse(user user.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
