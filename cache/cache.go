package cache

import (
	"awesomeProject/internal/model"
	"errors"
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	cache map[string]*model.Order
}

func NewCache() *Cache {
	return &Cache{cache: make(map[string]*model.Order)}
}

func (c *Cache) Put(key string, data *model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = data
}

func (c *Cache) GetByIdFromCache(key string) (*model.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.cache[key]
	if found != true {
		return nil, errors.New("no data in cache by key")
	}
	return value, nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, found := c.cache[key]; found != true {
		return errors.New("key does not exist")
	}
	delete(c.cache, key)
	return nil
}
