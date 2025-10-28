package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/adapters/cache"
	httpTransport "github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/adapters/http/handlers"
	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/adapters/repository/inmemory"
	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/user"
)

func main() {

	/* cache, err := cache.InitRistrettoCache()
	if err != nil {
		log.Fatal(err)
	} */
	userCache := cache.InitLRUCacheWithEvict[string, *user.User](1000)
	var userRepo user.UserRepository = inmemory.NewUserRepository(userCache)

	// Initialize your UserService with a repository implementation
	userService := user.NewUserService(userRepo)
	userHandler := httpTransport.NewUserHandler(*userService)

	// Create chi router
	r := chi.NewRouter()

	//Routes
	// User Routes
	r.Get("/users", userHandler.List)
	r.Post("/users", userHandler.Create)
	r.Get("/users/{id}", userHandler.Get)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
