package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserFavouritesResponse struct {
	AssetType   string `json:"assetType"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
