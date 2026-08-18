package proxy

import (
	"sync/atomic"
)

type RoundRobinBalancer struct {
	backends []string
	counter  uint64
}

func NewRoundRobinBalancer(backends []string) *RoundRobinBalancer {
	return &RoundRobinBalancer{backends: backends}
}

func (b *RoundRobinBalancer) Next() string {
	if len(b.backends) == 0 {
		return ""
	}
	idx := atomic.AddUint64(&b.counter, 1) % uint64(len(b.backends))
	return b.backends[idx]
}
