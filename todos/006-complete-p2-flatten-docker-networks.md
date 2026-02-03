---
status: complete
priority: p2
issue_id: "006"
tags: [simplicity, docker, devops]
dependencies: []
---

# Problem Statement
PR #426 introduced dual-network segmentation (frontend/backend). For a single-node Docker Compose setup where the application bridges both networks, this provides minimal security gain while adding architectural noise and making maintenance harder.

# Findings
- `deploy/docker-compose.yml` and tunnel files use separate networks.
- Both DHH and Code Simplicity agents recommend flattening this for "Omakase" simplicity.

# Proposed Solutions

## Option 1: Flatten to single sub2api-network
Remove `frontend-network` and `backend-network`. Move all containers to a single `sub2api-network`.
- **Pros**: Simplifies configuration, follows YAGNI, reduces cognitive load.
- **Cons**: Minor reduction in theoretically "ideal" network segmentation.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] Docker Compose files updated to use a single network.
- [x] Connectivity verified between app and DB/Redis.
- [x] Documentation updated.

# Work Log
## 2026-02-03 - Initial Finding
Decision made to prioritize simplicity over "infrastructure theatre" based on expert agent reviews.

## 2026-02-03 - Resolved
Flattened networks in all compose files. Updated documentation.
