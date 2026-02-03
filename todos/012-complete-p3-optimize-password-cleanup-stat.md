---
status: complete
priority: p3
issue_id: "012"
tags: [performance, ops]
dependencies: []
---

# Problem Statement
`CleanupInitialPassword()` performs an `os.Stat` on every successful login, which is an unnecessary filesystem operation.

# Findings
- `backend/internal/setup/setup.go` check-then-act pattern.
- In high-load scenarios with automated logins, this adds negligible but avoidable overhead.

# Proposed Solutions

## Option 1: sync.Once or direct remove
Call `os.Remove` directly and ignore `os.ErrNotExist`, or use a package-level variable to skip the check after the first success.
- **Pros**: Cleaner logic, minor performance win.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] `CleanupInitialPassword` optimized.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Performance Oracle.

## 2026-02-03 - Resolution
Optimized `CleanupInitialPassword` by removing `os.Stat` and using `atomic.Bool` to avoid unnecessary filesystem operations on every login.
