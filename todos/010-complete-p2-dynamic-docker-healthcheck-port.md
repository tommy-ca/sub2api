---
status: complete
priority: p2
issue_id: "010"
tags: [docker, ops]
dependencies: []
---

# Problem Statement
The Docker Compose healthcheck for the `sub2api` service uses a hardcoded port (8080), which may conflict with user-defined `SERVER_PORT`.

# Findings
- `deploy/docker-compose.yml` uses `http://localhost:8080/health`.
- The `Dockerfile` uses `${SERVER_PORT}` for its internal healthcheck.
- If a user changes `SERVER_PORT` in `.env`, the Compose healthcheck will fail even if the app is healthy.

# Proposed Solutions

## Option 1: Use variable in compose
Update `docker-compose.yml` to use `${SERVER_PORT:-8080}` in the healthcheck URL.
- **Pros**: Accuracy across configurations.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [ ] `docker-compose.yml` healthcheck uses dynamic port.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Pattern Recognition Specialist during PR #426 review.
