package ports

import (
	"net/http"
)

type UserHandler interface {
	// Create handles HTTP POST /users requests
	Create(w http.ResponseWriter, r *http.Request)

	// Get handles HTTP GET /users/{id} requests
	Get(w http.ResponseWriter, r *http.Request)

	// List handles HTTP GET /users requests
	List(w http.ResponseWriter, r *http.Request)

	// Update handles HTTP PUT /users/{id} requests
	Update(w http.ResponseWriter, r *http.Request)

	// Delete handles HTTP DELETE /users/{id} requests
	Delete(w http.ResponseWriter, r *http.Request)

	// Get handles HTTP GET /users/{id}/favourites requests
	GetFavourites(w http.ResponseWriter, r *http.Request)
}

type FavouriteHandler interface {
	// Create handles HTTP POST /favourites requests
	Create(w http.ResponseWriter, r *http.Request)

	// Delete handles HTTP DELETE /favourites/{id} requests
	Delete(w http.ResponseWriter, r *http.Request)
}

type AudienceHandler interface {
	// Create handles HTTP POST /audiences requests
	Create(w http.ResponseWriter, r *http.Request)

	// Get handles HTTP GET /audiences/{id} requests
	Get(w http.ResponseWriter, r *http.Request)

	// List handles HTTP GET /audiences requests
	List(w http.ResponseWriter, r *http.Request)

	// Update handles HTTP PUT /audiences/{id} requests
	Update(w http.ResponseWriter, r *http.Request)

	// Delete handles HTTP DELETE /audiences/{id} requests
	Delete(w http.ResponseWriter, r *http.Request)
}

type InsightHandler interface {
	// Create handles HTTP POST /insights requests
	Create(w http.ResponseWriter, r *http.Request)

	// Get handles HTTP GET /insights/{id} requests
	Get(w http.ResponseWriter, r *http.Request)

	// List handles HTTP GET /insights requests
	List(w http.ResponseWriter, r *http.Request)

	// Update handles HTTP PUT /insights/{id} requests
	Update(w http.ResponseWriter, r *http.Request)

	// Delete handles HTTP DELETE /insights/{id} requests
	Delete(w http.ResponseWriter, r *http.Request)
}

type ChartHandler interface {
	// Create handles HTTP POST /charts requests
	Create(w http.ResponseWriter, r *http.Request)

	// Get handles HTTP GET /charts/{id} requests
	Get(w http.ResponseWriter, r *http.Request)

	// List handles HTTP GET /charts requests
	List(w http.ResponseWriter, r *http.Request)

	// Update handles HTTP PUT /charts/{id} requests
	Update(w http.ResponseWriter, r *http.Request)

	// Delete handles HTTP DELETE /charts/{id} requests
	Delete(w http.ResponseWriter, r *http.Request)
}
