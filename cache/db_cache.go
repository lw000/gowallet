package cache

import (
	"github.com/lw000/gocommon/cache"
	"sync"
)

var (
	commonCache     tycache.MemCache
	commonCacheOnce sync.Once
)

func CommonCacheService() tycache.MemCache {
	commonCacheOnce.Do(func() {
		commonCache = tycache.NewLRU(1024 * 100)
	})
	return commonCache
}
