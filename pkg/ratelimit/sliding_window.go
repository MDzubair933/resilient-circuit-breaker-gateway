package ratelimit

import (
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	mu           sync.Mutex
	windowSize   time.Duration
	limit        int
	requestTimes []time.Time
}

func NewSlidingWindowLimiter(limit int, window time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:        limit,
		windowSize:   window,
		requestTimes: make([]time.Time, 0, limit),
	}
}

func (s *SlidingWindowLimiter) Allow() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-s.windowSize)

	validIdx := 0
	for _, t := range s.requestTimes {
		if t.After(cutoff) {
			break
		}
		validIdx++
	}
	s.requestTimes = s.requestTimes[validIdx:]

	if len(s.requestTimes) < s.limit {
		s.requestTimes = append(s.requestTimes, now)
		return true
	}
	return false
}
