package circuit

import "sync"

type FailureMetrics struct {
	mu            sync.Mutex
	totalRequests uint64
	failedRequests uint64
}

func (m *FailureMetrics) Record(success bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.totalRequests++
	if !success {
		m.failedRequests++
	}
}

func (m *FailureMetrics) FailureRate() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.totalRequests == 0 {
		return 0.0
	}
	return float64(m.failedRequests) / float64(m.totalRequests)
}
