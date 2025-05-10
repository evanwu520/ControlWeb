package cache

import (
	"sync"
	"time"
)

type adminInfo struct {
	Account   string
	CreatedAt int64
}

type tokenCache struct {
	mu    sync.RWMutex
	store map[string]adminInfo
}

var adminCache = tokenCache{
	store: make(map[string]adminInfo),
}

// Set sets the value for a given key
func SetToken(key, account string) {
	adminCache.mu.Lock()
	defer adminCache.mu.Unlock()
	adminCache.store[key] = adminInfo{Account: account, CreatedAt: time.Now().Unix()}
}

// Get retrieves the value for a given key
func GetToken(key string) (adminInfo, bool) {
	adminCache.mu.RLock()
	defer adminCache.mu.RUnlock()
	val, ok := adminCache.store[key]
	return val, ok
}

// Delete removes the key from the cache
func Delete(key string) {
	adminCache.mu.Lock()
	defer adminCache.mu.Unlock()
	delete(adminCache.store, key)
}

// Clear resets the cache
func ClearToken() {
	adminCache.mu.Lock()
	defer adminCache.mu.Unlock()
	adminCache.store = make(map[string]adminInfo)
}
