---
status: complete
priority: p2
issue_id: "014"
tags: [security, refactor, quality]
dependencies: []
---

# Problem Statement
There is duplicated and inconsistent logic for IP/Host blocking. `validator.go` has `isBlockedHost` which is less comprehensive than the hardened `IsPrivateIP`.

# Findings
- `backend/internal/util/urlvalidator/validator.go`: `isBlockedHost` manually checks loopback/private but misses the extra ranges (CGNAT, Benchmarking, Teredo) found in `IsPrivateIP`.
- Code Simplicity agent flagged this redundancy.

# Proposed Solutions

## Option 1: Unify in isBlockedHost (Recommended)
Update `isBlockedHost` to use `IsPrivateIP(net.ParseIP(host))` after the "localhost" string check.
- **Pros**: Single source of truth for blocked ranges.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] `isBlockedHost` uses `IsPrivateIP` logic.
- [x] SSRF protection is consistently applied across all validation paths.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Code Simplicity agent.

## 2026-02-04 - Implementation
Updated `isBlockedHost` to use `IsPrivateIP` for consistent IP blocking logic. Verified with existing tests.
