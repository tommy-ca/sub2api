---
status: complete
priority: p1
issue_id: "013"
tags: [security, reliability, redis]
dependencies: []
---

# Problem Statement
Per `CLAUDE.md`, the application should fail at boot if it cannot reach critical dependencies (Postgres, Redis). Currently, `InitRedis` and `ProvideRedis` do not perform a `Ping()` check, allowing the app to start even if Redis is down.

# Findings
- `backend/internal/repository/redis.go`: `InitRedis` creates the client but doesn't verify connectivity.
- Kieran flagged this as a "Fail Fast" violation.

# Proposed Solutions

## Option 1: Add Ping to InitRedis (Recommended)
Add a `Ping(ctx)` call to `InitRedis` and `log.Fatalf` if it fails.
- **Pros**: Strict compliance with testing philosophy.
- **Cons**: App won't start if Redis is transiently down during boot.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] Application fails to start if Redis is unreachable during boot.
- [x] Error message clearly indicates Redis connection failure.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Kieran during PR review.

## 2026-02-04 - Resolution
Added `Ping(ctx)` call to `InitRedis` in `backend/internal/repository/redis.go` with `log.Fatalf` on failure. This ensures the application fails fast if Redis is unreachable at boot.
Note: The todo previously mentioned `backend/internal/setup/setup.go`, but the function was found in `backend/internal/repository/redis.go`.