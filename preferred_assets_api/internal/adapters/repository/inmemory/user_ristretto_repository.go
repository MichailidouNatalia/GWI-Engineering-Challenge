package inmemory

/*
import (
	"encoding/json"
	"errors"

	"github.com/dgraph-io/ristretto"

	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/user"
)

type RistrettoUserRepository struct {
	cache *ristretto.Cache
}

func NewUserRepository(cache *ristretto.Cache) *RistrettoUserRepository {
	return &RistrettoUserRepository{cache: cache}
}

func (r *RistrettoUserRepository) Save(user user.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	r.cache.Set(user.ID, data, 1)
	r.cache.Wait()
	return nil
}

func (r *RistrettoUserRepository) GetByID(id string) (user.User, error) {
	value, found := r.cache.Get(id)
	if !found {
		return user.User{}, errors.New("user not found")
	}
	data, ok := value.([]byte)
	if !ok {
		return user.User{}, errors.New("invalid cache type")
	}
	var usr user.User
	if err := json.Unmarshal(data, &usr); err != nil {
		return user.User{}, err
	}
	return usr, nil
}

func (r *RistrettoUserRepository) GetAll() ([]user.User, error) {
	return nil, errors.New("GetAll not implemented: ristretto has no iteration")
}

func (r *RistrettoUserRepository) Delete(id string) error {
	r.cache.Del(id)
	return nil
}
*/
