package server

import (
	"log"
	"net/http"
	"time"

	_ "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/docs"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/cmd/api/config"
	httpTransport "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/handlers"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/http/middleware"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/inmemory"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/dto"
	application "github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/application/services"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/pkg/auth"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/pkg/cache"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	UserHandler      *httpTransport.UserHandler
	FavouriteHandler *httpTransport.FavouriteHandler
	Keycloak         *auth.KeycloakClient
	Config           *config.Config
}

func New() *App {
	cfg := config.Load()

	// Initialize Keycloak client
	keycloakClient, _ := auth.NewKeycloakClient(&cfg.Keycloak)

	//Initialization for User resources
	userCache := cache.InitLRUCacheWithEvict[string, *entities.UserEntity](3)
	assetCache := cache.InitLRUCacheWithEvict[string, entities.AssetEntity](50)
	var userRepo ports.UserRepository = inmemory.NewUserRepository(userCache)
	var assetRepo ports.AssetRepository = inmemory.NewAssetRepository(assetCache)
	userService := application.NewUserService(userRepo, assetRepo)
	userHandler := httpTransport.NewUserHandler(*userService)

	//Initialization for Favourite resources
	userFavouritesCache := cache.InitLRUCacheWithEvict[string, map[string]time.Time](3)
	favouriteExistsCache := cache.InitLRUCacheWithEvict[string, bool](9)
	var favouriteRepo ports.FavouriteRepository = inmemory.NewFavouriteRepository(userFavouritesCache, favouriteExistsCache)
	favouriteService := application.NewFavouriteService(favouriteRepo)
	favouriteHandler := httpTransport.NewFavouriteHandler(*favouriteService)

	return &App{
		UserHandler:      userHandler,
		FavouriteHandler: favouriteHandler,
		Keycloak:         keycloakClient,
		Config:           cfg,
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

		//Group Users
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
		apiRouter.With(middleware.RequireAnyRole("Users")).
			Get("/users/{id}/favourites", application.UserHandler.GetFavourites)

		//Group Favourites
		apiRouter.With(middleware.RequireAnyRole("Users")).With(middleware.ValidateBody[dto.FavouriteRequest]()).
			Post("/favourites", application.FavouriteHandler.Create)
		apiRouter.With(middleware.RequireAnyRole("Users")).
			Post("/favourites/{userId}/assets/{assetId}", application.FavouriteHandler.Delete)
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("🚀 Server running on :8081")
	log.Println("📘 Swagger docs at http://localhost:8081/swagger/index.html")

	return http.ListenAndServe(":8081", router)
}
