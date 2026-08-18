package proxy

import (
	"net/http"
	"time"
)

type HealthChecker struct {
	TargetURL string
	Interval  time.Duration
	IsHealthy bool
}

func (h *HealthChecker) Probe() bool {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(h.TargetURL)
	h.IsHealthy = (err == nil && resp.StatusCode == 200)
	return h.IsHealthy
}
