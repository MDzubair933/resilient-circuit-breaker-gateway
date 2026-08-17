package ratelimit

import "context"

type DistributedLimiter interface {
	Allow(ctx context.Context, key string, limit int) (bool, error)
}

type MockRedisLimiter struct{}

func (m *MockRedisLimiter) Allow(ctx context.Context, key string, limit int) (bool, error) {
	// Distributed Lua script evaluation
	return true, nil
}
