---
status: complete
priority: p2
issue_id: "011"
tags: [refactor, quality]
dependencies: []
---

# Problem Statement
Several files still use manual `os.Getenv` and string comparisons for boolean flags instead of the new `envutil` package.

# Findings
- Manual parsing found in `antigravity_gateway_service.go`, `gateway_service.go`, and `ops_ws_handler.go`.
- Inconsistent behavior (e.g., some check for "on", some for "true").

# Proposed Solutions

## Option 1: Complete Migration
Migrate all environment-based boolean and integer parsing to `envutil`.
- **Pros**: Consistency, reduced duplication.
- **Cons**: None.
- **Effort**: Medium
- **Risk**: Low

# Acceptance Criteria
- [x] No manual `os.Getenv` + string checks for booleans in core services.
- [x] `envutil` handles "on/off" to support existing patterns.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Pattern Recognition Specialist.

## 2026-02-03 - Completed
Migrated manual parsing in `antigravity_gateway_service.go`, `gateway_service.go`, and `ops_ws_handler.go` to use `envutil`.
Removed unused `os` and `strconv` imports.
