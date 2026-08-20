package tests

import (
	"errors"
	"testing"
	"time"
	"github.com/MDzubair933/resilient-circuit-breaker-gateway/pkg/circuit"
)

func TestCircuitBreakerTrip(t *testing.T) {
	cb := circuit.NewCircuitBreaker(2, 50*time.Millisecond)
	errMock := errors.New("upstream timeout")

	cb.Execute(func() error { return errMock })
	cb.Execute(func() error { return errMock })

	// Third call should fail fast
	err := cb.Execute(func() error { return nil })
	if err == nil {
		t.Fatal("expected circuit breaker to fast-fail open")
	}
}
