package application

import (
	"fmt"
	"log"
	"time"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/entities"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/adapters/repositories/mapper"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
	"golang.org/x/crypto/bcrypt"
)

var _ ports.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repo      ports.UserRepository
	assetRepo ports.AssetRepository
}

func NewUserService(usrRepo ports.UserRepository, assetRepo ports.AssetRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: usrRepo,
		assetRepo: assetRepo}
}

func (usrService UserServiceImpl) GetUserByID(id string) (*domain.User, error) {
	user, err := usrService.repo.GetByID(id)
	return mapper.UserEntityToDomain(user), err
}

func (usrService UserServiceImpl) GetAllUsers() ([]domain.User, error) {
	users, err := usrService.repo.GetAll()
	userList := mapper.UserEntintyToDomainList(users)
	return userList, err
}

func (usrService UserServiceImpl) CreateUser(usr domain.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(hashedPassword)
	usr.CreatedAt = time.Now().UTC()
	user := mapper.UserEntityFromDomain(usr)

	return usrService.repo.Save(user)
}

func (usrService UserServiceImpl) DeleteUser(id string) error {
	return usrService.repo.Delete(id)
}

func (usrService UserServiceImpl) UpdateUser(usr domain.User) error {
	usr.UpdatedAt = time.Now().UTC()
	user := mapper.UserEntityFromDomain(usr)
	return usrService.repo.Update(user)
}

func (usrService UserServiceImpl) GetFavouritesByUser(id string) ([]domain.Favourite, error) {
	favs, err := usrService.repo.GetFavouritesByID(id)
	favsList := mapper.FavouriteEntityToDomainList(favs)
	enhancedFavs, err := usrService.batchEnhanceFavourites(favsList)
	if err != nil {
		return nil, err
	}

	return enhancedFavs, err
}

func (usrService UserServiceImpl) batchEnhanceFavourites(favourites []domain.Favourite) ([]domain.Favourite, error) {
	if len(favourites) == 0 {
		return favourites, nil
	}

	assetIDs := make([]string, 0, len(favourites))
	for _, fav := range favourites {
		assetIDs = append(assetIDs, fav.AssetID)
	}

	// Batch fetch all assets
	assetsEntities, err := usrService.assetRepo.GetByIDs(assetIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch assets: %w", err)
	}

	// Create a map for quick lookup
	assetMap := make(map[string]entities.AssetEntity)
	for _, assetEntity := range assetsEntities {
		assetMap[assetEntity.GetID()] = assetEntity
	}

	// Enhance favourites with assets
	enhancedFavs := make([]domain.Favourite, 0, len(favourites))
	for _, fav := range favourites {
		assetEntity, exists := assetMap[fav.AssetID]
		if !exists {
			log.Printf("Asset %s not found for favourite", fav.AssetID)
			continue
		}

		// Map Asset Entities to Domain
		asset, err := mapper.AssetEntityToDomain(assetEntity)
		if err != nil {
			log.Printf("Failed to map asset entity: %v", err)
			continue
		}
		// Assign Asset domain objects to Favourite domain object
		if err := fav.SetAsset(asset); err != nil {
			log.Printf("Failed to set asset on favourite: %v", err)
			continue
		}

		enhancedFavs = append(enhancedFavs, fav)
	}

	return enhancedFavs, nil
}
