---
status: complete
priority: p2
issue_id: "015"
tags: [refactor, quality]
dependencies: []
---

# Problem Statement
While `envutil` was introduced, some core environment variables are still being parsed manually using `os.Getenv`, missing out on the standardized trimming and default handling.

# Findings
- `DATA_DIR` in `setup.go` and `config.go` uses manual parsing.
- `TZ` in `usage_log_repo.go` uses manual parsing.
- Pattern Recognition Specialist flagged these as inconsistent.

# Proposed Solutions

## Option 1: Complete Migration to envutil (Recommended)
Replace all remaining `os.Getenv` + string checks with `envutil` helpers.
- **Pros**: Perfect consistency.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] `DATA_DIR` and `TZ` migrated to `envutil.GetString`.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Pattern Recognition Specialist.

## 2026-02-04 - Resolved
Migrated `DATA_DIR` and `TZ` usage to `envutil.GetString` in `setup.go`, `config.go`, and `usage_log_repo.go`.
