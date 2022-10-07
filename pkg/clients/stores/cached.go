package stores

import (
	"context"
	"time"

	"github.com/Tlantic/go-util/v4/cache"
	"github.com/Tlantic/go-util/v4/mrs"
	"google.golang.org/grpc"
)

type cachedTransport struct {
	transport
	cache *cache.Map
	ttl   time.Duration
}

var _ transport = (*cachedTransport)(nil)

func (c *cachedTransport) GetStore(ctx context.Context, storeId string, opts ...grpc.CallOption) (*Store, error) {
	key := mrs.GetMRSContext(ctx).Organization.Code + ":" + storeId
	if s, _ := c.cache.Get(key).(*Store); s != nil {
		return s, nil
	}

	s, err := c.transport.GetStore(ctx, storeId, opts...)
	if err != nil {
		return nil, err
	}

	c.cache.Set(key, s, c.ttl)
	return s, nil
}
