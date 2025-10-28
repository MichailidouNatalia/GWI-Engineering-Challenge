package inmemory

import (
	"errors"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/domain/user"
	"github.com/MichailidouNatalia/GWI-Engineering-Challenge/preferred_assets_api/internal/ports"
)

var _ ports.UserRepository = (*LRUUserRepositoryImpl)(nil)

type LRUUserRepositoryImpl struct {
	cache *lru.Cache[string, *user.User]
	mu    sync.RWMutex
}

func NewUserRepository(cache *lru.Cache[string, *user.User]) *LRUUserRepositoryImpl {
	return &LRUUserRepositoryImpl{cache: cache}
}

func (r *LRUUserRepositoryImpl) Save(u user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache.Add(u.ID, &u)
	return nil
}

func (r *LRUUserRepositoryImpl) GetByID(id string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.cache.Get(id)
	if !ok {
		return &user.User{}, errors.New("user not found")
	}
	return val, nil
}

func (r *LRUUserRepositoryImpl) GetAll() ([]user.User, error) {
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

func (r *LRUUserRepositoryImpl) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache.Remove(id)
	return nil
}

func (r *LRUUserRepositoryImpl) Update(u user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.cache.Get(u.ID); !ok {
		return errors.New("user not found")
	}

	r.cache.Add(u.ID, &u) // Overwrite existing
	return nil
}
