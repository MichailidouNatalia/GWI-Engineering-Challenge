# preferred_assets_api — Project Structure (generated)

Top-level layout and key files for the preferred_assets_api workspace. Use the links to open files.

- [Dockerfile](preferred_assets_api/Dockerfile)
- [docker-compose.yml](preferred_assets_api/docker-compose.yml)
- [.env/](preferred_assets_api/.env/)
  - [env.file](preferred_assets_api/.env/env.file)
- [go.mod](preferred_assets_api/go.mod)
- [go.sum](preferred_assets_api/go.sum)
- [README.md](../README.md) (project root)

preferred_assets_api/
- cmd/
  - api/
    - config/
      - [config.go](preferred_assets_api/cmd/api/config/config.go)
      - [keycloak-realm-import.json](preferred_assets_api/cmd/api/config/keycloak-realm-import.json)
      - [users-export.json](preferred_assets_api/cmd/api/config/users-export.json)
    - main/
      - [main.go](preferred_assets_api/cmd/api/main/main.go)
    - server/
      - [server.go](preferred_assets_api/cmd/api/server/server.go) — [`server.New`](preferred_assets_api/cmd/api/server/server.go), [`server.App.Run`](preferred_assets_api/cmd/api/server/server.go)
- docs/
  - [docs.go](preferred_assets_api/docs/docs.go)
  - [swagger.yaml](preferred_assets_api/docs/swagger.yaml)
  - [swagger.json](preferred_assets_api/docs/swagger.json)
- internal/
  - adapters/
    - http/
      - handlers/
        - [user_handler.go](preferred_assets_api/internal/adapters/http/handlers/user_handler.go) — [`handlers.UserHandler`](preferred_assets_api/internal/adapters/http/handlers/user_handler.go)
        - [asset_handler.go](preferred_assets_api/internal/adapters/http/handlers/asset_handler.go)
        - [favourite_handler.go](preferred_assets_api/internal/adapters/http/handlers/favourite_handler.go)
      - middleware/
        - [auth.go](preferred_assets_api/internal/adapters/http/middleware/auth.go) — [`middleware.AuthMiddleware`](preferred_assets_api/internal/adapters/http/middleware/auth.go), [`middleware.RoleMiddleware`](preferred_assets_api/internal/adapters/http/middleware/auth.go), [`middleware.RequireAnyRole`](preferred_assets_api/internal/adapters/http/middleware/auth.go)
        - [requests_validator.go](preferred_assets_api/internal/adapters/http/middleware/requests_validator.go) — [`middleware.ValidateBody`](preferred_assets_api/internal/adapters/http/middleware/requests_validator.go), [`middleware.GetValidatedBody`](preferred_assets_api/internal/adapters/http/middleware/requests_validator.go)
        - [validation_error.go](preferred_assets_api/internal/adapters/http/middleware/validation_error.go)
  - repositories/
    - entities/
      - [asset_entity.go](preferred_assets_api/internal/adapters/repositories/entities/asset_entity.go)
      - [asset_chart_entity.go](preferred_assets_api/internal/adapters/repositories/entities/asset_chart_entity.go)
      - [asset_audience_entity.go](preferred_assets_api/internal/adapters/repositories/entities/asset_audience_entity.go)
      - [asset_insight_entity.go](preferred_assets_api/internal/adapters/repositories/entities/asset_insight_entity.go)
      - [asset_types_entity.go](preferred_assets_api/internal/adapters/repositories/entities/asset_types_entity.go)
      - [favourite_entity.go](preferred_assets_api/internal/adapters/repositories/entities/favourite_entity.go)
      - [user_entity.go](preferred_assets_api/internal/adapters/repositories/entities/user_entity.go)
    - inmemory/
      - [user_lru_repository.go](preferred_assets_api/internal/adapters/repositories/inmemory/user_lru_repository.go) — LRU user repo, uses [`cache.InitLRUCacheWithEvict`](preferred_assets_api/pkg/cache/lru_cache.go)
      - [favourite_lru_repository.go](preferred_assets_api/internal/adapters/repositories/inmemory/favourite_lru_repository.go)
      - [asset_lru_repository.go](preferred_assets_api/internal/adapters/repositories/inmemory/asset_lru_repository.go)
      - [asset_lru_repository_test.go](preferred_assets_api/internal/adapters/repositories/inmemory/asset_lru_repository_test.go)
    - mapper/
      - [user_entity_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/user_entity_mapper.go)
      - [asset_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/asset_mapper.go)
      - [asset_chart_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/asset_chart_mapper.go)
      - [asset_audience_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/asset_audience_mapper.go)
      - [asset_insight_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/asset_insight_mapper.go)
      - [favourite_entity_mapper.go](preferred_assets_api/internal/adapters/repositories/mapper/favourite_entity_mapper.go)
    - persistence/ (empty / DB adapters)
  - application/
    - dto/
      - [user_dto.go](preferred_assets_api/internal/application/dto/user_dto.go)
      - [asset_dto.go](preferred_assets_api/internal/application/dto/asset_dto.go)
      - [audience_dto.go](preferred_assets_api/internal/application/dto/audience_dto.go)
      - [chart_dto.go](preferred_assets_api/internal/application/dto/chart_dto.go)
      - [insight_dto.go](preferred_assets_api/internal/application/dto/insight_dto.go)
      - [favourite_dto.go](preferred_assets_api/internal/application/dto/favourite_dto.go)
    - mapping/
      - [user_mapper.go](preferred_assets_api/internal/application/mapping/user_mapper.go) — [`mapping.ToDomain`](preferred_assets_api/internal/application/mapping/user_mapper.go), [`mapping.UpdateToDomain`](preferred_assets_api/internal/application/mapping/user_mapper.go)
      - [asset_mapper.go](preferred_assets_api/internal/application/mapping/asset_mapper.go)
      - [favourite_mapper.go](preferred_assets_api/internal/application/mapping/favourite_mapper.go)
    - services/
      - [user_service.go](preferred_assets_api/internal/application/services/user_service.go) — [`application.NewUserService`](preferred_assets_api/internal/application/services/user_service.go), [`application.UserServiceImpl`](preferred_assets_api/internal/application/services/user_service.go)
      - [favourite_service.go](preferred_assets_api/internal/application/services/favourite_service.go)
      - [asset_service.go](preferred_assets_api/internal/application/services/asset_service.go)
      - [asset_audience_service.go](preferred_assets_api/internal/application/services/asset_audience_service.go)
      - [asset_chart_service.go](preferred_assets_api/internal/application/services/asset_chart_service.go)
      - [asset_insight_service.go](preferred_assets_api/internal/application/services/asset_insight_service.go)
  - domain/
    - [user.go](preferred_assets_api/internal/domain/user.go)
    - [favourite.go](preferred_assets_api/internal/domain/favourite.go)
    - [asset.go](preferred_assets_api/internal/domain/asset.go)
    - [asset_type.go](preferred_assets_api/internal/domain/asset_type.go)
    - [asset_chart.go](preferred_assets_api/internal/domain/asset_chart.go)
    - [asset_audience.go](preferred_assets_api/internal/domain/asset_audience.go)
    - [asset_insight.go](preferred_assets_api/internal/domain/asset_insight.go)
    - errors/ (domain errors)
  - ports/
    - [http.go](preferred_assets_api/internal/ports/http.go) — HTTP handler interface
    - [repositories.go](preferred_assets_api/internal/ports/repositories.go) — repository interfaces (UserRepository, FavouriteRepository, ChartRepository, InsightRepository, AudienceRepository)
    - [services.go](preferred_assets_api/internal/ports/services.go) — [`ports.UserService`](preferred_assets_api/internal/ports/services.go)
- pkg/
  - auth/
    - [keycloak.go](preferred_assets_api/pkg/auth/keycloak.go) — [`auth.NewKeycloakClient`](preferred_assets_api/pkg/auth/keycloak.go), [`auth.KeycloakClient.VerifyToken`](preferred_assets_api/pkg/auth/keycloak.go), [`auth.KeycloakClient.IntrospectToken`](preferred_assets_api/pkg/auth/keycloak.go), [`auth.KeycloakClient.GetUserRoles`](preferred_assets_api/pkg/auth/keycloak.go)
  - cache/
    - [lru_cache.go](preferred_assets_api/pkg/cache/lru_cache.go) — [`cache.InitLRUCache`](preferred_assets_api/pkg/cache/lru_cache.go), [`cache.InitLRUCacheWithEvict`](preferred_assets_api/pkg/cache/lru_cache.go)
  - logger/ (logging helpers)

Notes
- The main server wiring lives in [cmd/api/server/server.go](preferred_assets_api/cmd/api/server/server.go) and is responsible for creating Keycloak client (`auth.NewKeycloakClient`), initializing caches (`cache.InitLRUCacheWithEvict`), repositories (`inmemory.NewUserRepository`), services (`application.NewUserService`) and handlers (`httpTransport.NewUserHandler`).
- Auth is implemented in [internal/adapters/http/middleware/auth.go](preferred_assets_api/internal/adapters/http/middleware/auth.go) and uses [`auth.KeycloakClient`](preferred_assets_api/pkg/auth/keycloak.go).
- LRU-based in-memory repositories are under [internal/adapters/repositories/inmemory](preferred_assets_api/internal/adapters/repositories/inmemory).
- API handlers and request validation middleware are under [internal/adapters/http](preferred_assets_api/internal/adapters/http).

How to use
- Build: `docker compose up --build` or `go build ./cmd/api/...`
- Run with Dockerfile: see [preferred_assets_api/Dockerfile](preferred_assets_api/Dockerfile)

If you want this saved to a different path or a different format (JSON, tree, or plantuml), tell me the target filepath.