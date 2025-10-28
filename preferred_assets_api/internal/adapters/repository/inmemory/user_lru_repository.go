package inmemory

import (
	"errors"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/MichailidouNatalia/preferred_assets_api/preferred_assets_api/internal/domain/user"
)

type LRUUserRepository struct {
	cache *lru.Cache[string, *user.User]
	mu    sync.RWMutex
}

func NewUserRepository(cache *lru.Cache[string, *user.User]) *LRUUserRepository {
	return &LRUUserRepository{cache: cache}
}

func (r *LRUUserRepository) Save(u user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache.Add(u.ID, &u)
	return nil
}

func (r *LRUUserRepository) GetByID(id string) (user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.cache.Get(id)
	if !ok {
		return user.User{}, errors.New("user not found")
	}
	return *val, nil
}

func (r *LRUUserRepository) GetAll() ([]user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	users := make([]user.User, 0, r.cache.Len())
	for _, key := range r.cache.Keys() {
		if val, ok := r.cache.Peek(key); ok {
			users = append(users, *val)
		}
	}
	return users, nil
}

func (r *LRUUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache.Remove(id)
	return nil
}
