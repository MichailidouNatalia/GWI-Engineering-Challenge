package main

import (
	"log"
	"os"

	app "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/cmd/api/server"
)

// @title           Preferred Assets API
// @version         1.0
// @description     API for managing users
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /
func main() {
	// Check if we're running in Docker
	if os.Getenv("DOCKER_ENV") == "true" {
		log.Println("Running in Docker container")
	}

	a := app.New()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

	/* userCache := cache.InitLRUCacheWithEvict[string, *user.User](3)
	var userRepo user.UserRepository = inmemory.NewUserRepository(userCache)

	// Initialize your UserService with a repository implementation
	userService := user.NewUserService(userRepo)
	userHandler := httpTransport.NewUserHandler(*userService)

	// Create chi router
	r := chi.NewRouter()

	//Routes
	// User Routes
	r.Get("/users", userHandler.List)
	r.Get("/users/{id}", userHandler.Get)
	r.With(middleware.ValidateBody[dto.CreateUserRequest]()).Post("/users", userHandler.Create)
	r.Put("/users/{id}", userHandler.Update)
	r.Delete("/users/{id}", userHandler.Get)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	} */
}
