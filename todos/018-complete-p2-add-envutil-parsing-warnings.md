---
status: complete
priority: p2
issue_id: "018"
tags: [quality, failure-modes]
dependencies: []
---

# Problem Statement
`envutil.GetInt` silently returns the default on a parsing error, which can mask configuration mistakes.

# Findings
- Kieran flagged this as a "lazy" failure mode.
- Misconfigured environment variables (e.g. `SERVER_PORT=abc`) won't be visible in logs.

# Proposed Solutions

## Option 1: Add Warning Logs
Add a `log.Printf` warning when a key exists but is not a valid integer.
- **Pros**: Better observability.
- **Cons**: Minor dependency on `log`.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] `envutil.GetInt` logs a warning for invalid values.
- [x] `envutil.GetBool` logs a warning for invalid values.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Kieran.

## 2026-02-04 - Implementation
Added `log.Printf` warnings to both `GetInt` and `GetBool` in `backend/internal/util/envutil/env.go`.
Updated functions to use `strings.TrimSpace` and check for empty strings before parsing.
Verified with tests.
