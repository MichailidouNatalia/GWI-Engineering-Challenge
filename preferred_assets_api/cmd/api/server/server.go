package server

import (
	"log"
	"net/http"

	_ "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/docs"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/cmd/api/config"
	httpTransport "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/handlers"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/middleware"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/inmemory"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	application "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/services"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/pkg/auth"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/pkg/cache"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	UserHandler *httpTransport.UserHandler
	Keycloak    *auth.KeycloakClient
	Config      *config.Config
}

func New() *App {
	cfg := config.Load()

	// Initialize Keycloak client
	keycloakClient, _ := auth.NewKeycloakClient(&cfg.Keycloak)
	userCache := cache.InitLRUCacheWithEvict[string, *user.User](3)
	var userRepo ports.UserRepository = inmemory.NewUserRepository(userCache)
	userService := application.NewUserService(userRepo)
	userHandler := httpTransport.NewUserHandler(*userService)

	return &App{
		UserHandler: userHandler,
		Keycloak:    keycloakClient,
		Config:      cfg,
	}
}

// Run configures the routes and starts the server
func (application *App) Run() error {

	// Generate swagger docs automatically in development
	//generateSwaggerDocs()

	router := chi.NewRouter()

	// API routes
	router.Route("/api/v1", func(apiRouter chi.Router) {
		// Authenticate all API routes
		apiRouter.Use(middleware.AuthMiddleware(application.Keycloak))

		apiRouter.With(middleware.RequireAnyRole("Administrators")).
			Get("/users", application.UserHandler.List)
		apiRouter.With(middleware.RequireAnyRole("Administrators")).
			Get("/users/{id}", application.UserHandler.Get)
		apiRouter.With(middleware.RequireAnyRole("Administrators")).With(middleware.ValidateBody[dto.CreateUserRequest]()).
			Post("/users", application.UserHandler.Create)
		apiRouter.With(middleware.RequireAnyRole("Administrators")).
			Put("/users/{id}", application.UserHandler.Update)
		apiRouter.With(middleware.RequireAnyRole("Administrators")).
			Delete("/users/{id}", application.UserHandler.Delete)
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("🚀 Server running on :8081")
	log.Println("📘 Swagger docs at http://localhost:8081/swagger/index.html")

	return http.ListenAndServe(":8081", router)
}
