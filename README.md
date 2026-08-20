# Resilient Circuit Breaker Gateway 🛡️⚡

A blazing-fast reverse-proxy edge gateway engineered in Go, featuring a 3-state finite state machine circuit breaker, lock-free sliding-window rate limiting, and zero-downtime health checking.

## Key Features
- **3-State Circuit Breaker:** CLOSED -> OPEN -> HALF_OPEN automatic fail-fast and healing.
- **Microsecond Rate Limiting:** In-memory sliding window and distributed Redis token-bucket.
- **Zero-Allocation Routing:** High-throughput reverse proxying with OpenTelemetry hooks.

Crafted by **Md Zubair** ([@MDzubair933](https://github.com/MDzubair933)) — **Zubair Systems**.
