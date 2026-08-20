package tests

import (
	"testing"
	"time"
	"github.com/MDzubair933/resilient-circuit-breaker-gateway/pkg/ratelimit"
)

func TestSlidingWindowLimiter(t *testing.T) {
	limiter := ratelimit.NewSlidingWindowLimiter(2, 100*time.Millisecond)
	if !limiter.Allow() || !limiter.Allow() {
		t.Fatal("expected first 2 requests to succeed")
	}
	if limiter.Allow() {
		t.Fatal("expected third request to be throttled")
	}
}
