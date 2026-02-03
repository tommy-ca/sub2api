---
status: complete
priority: p2
issue_id: "004"
tags: [quality, refactor, go]
dependencies: []
---

# Problem Statement
PR #426 introduced a robust `getEnvBoolOrDefault` helper, but older parts of the codebase still use manual string comparisons (e.g., `== "true"`) for boolean environment variables.

# Findings
- New helper: `getEnvBoolOrDefault` in `backend/internal/setup/setup.go`.
- Old pattern usage found in `REDIS_ENABLE_TLS` check.
- Inconsistent parsing makes the configuration system less resilient to varied boolean formats (1/0, yes/no).

# Proposed Solutions

## Option 1: Migrate to existing envutil.GetBool
Identify all boolean environment variable parsing and migrate them to the existing `backend/internal/util/envutil/env.go:GetBool` helper.
- **Pros**: Consistency, better robustness, no redundant code.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] All boolean env checks in `setup.go` and `config.go` use the helper.
- [x] Unit tests for `getEnvBoolOrDefault` to ensure parity with standard Docker boolean formats.

# Work Log
## 2026-02-03 - Initial Finding
Detected by pattern-recognition-specialist agent.

## 2026-02-03 - Resolved
Verified `setup.go` and `config.go` are using `envutil.GetBool` consistently. Redundant helpers removed.
