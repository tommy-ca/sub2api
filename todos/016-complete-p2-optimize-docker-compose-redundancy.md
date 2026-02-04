---
status: complete
priority: p2
issue_id: "016"
tags: [docker, simplicity, maintenance]
dependencies: []
---

# Problem Statement
`docker-compose.local.yml` is a 220-line near-duplicate of `docker-compose.yml`. This creates a heavy maintenance burden.

# Findings
- The only difference is bind mounts vs named volumes.
- Code Simplicity agent flagged this as unnecessary complexity.

# Proposed Solutions

## Option 1: Use minimal override (Recommended)
Replace `docker-compose.local.yml` with a small `docker-compose.override.yml` that only contains the volume sections.
- **Pros**: Reduces YAML code by 90%; easier maintenance.
- **Cons**: Requires users to understand how multiple compose files work (already documented).
- **Effort**: Medium
- **Risk**: Low

# Acceptance Criteria
- [x] `docker-compose.local.yml` replaced with a minimal override pattern.
- [x] Deployment documentation updated.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Code Simplicity agent.

## 2026-02-04 - Implementation
- Created `deploy/docker-compose.override.yml` with bind mount overrides.
- Removed `deploy/docker-compose.local.yml`.
- Updated `deploy/docker-deploy.sh` to use the override pattern.
- Updated `README.md` and `README_CN.md` with new instructions.
