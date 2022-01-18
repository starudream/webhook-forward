package webhookForward

import (
	"time"

	"github.com/go-sdk/lib/cache"
	"github.com/go-sdk/lib/rdx"
)

var mc cache.Cache

func init() {
	if rdx.Default() == nil {
		mc = cache.NewMemoryCacheWithCleaner(72*time.Hour, time.Second, nil)
	} else {
		mc = cache.NewRedisCache(72*time.Hour, rdx.Default())
	}
}
