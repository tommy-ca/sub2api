---
status: complete
priority: p1
issue_id: "008"
tags: [security, networking]
dependencies: []
---

# Problem Statement
The SSRF blocklist in `IsPrivateIP` is missing a check for general multicast addresses.

# Findings
- `backend/internal/util/urlvalidator/validator.go` checks for loopback, private, and link-local multicast, but not general multicast (`224.0.0.0/4` and `ff00::/8`).
- An attacker could potentially use multicast addresses to probe the internal network or cause unexpected behavior.

# Proposed Solutions

## Option 1: Use ip.IsMulticast()
Update `IsPrivateIP` to include a call to `ip.IsMulticast()`.
- **Pros**: Comprehensive coverage.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [ ] `IsPrivateIP` includes `ip.IsMulticast()` check.
- [ ] Tests verify that multicast addresses are blocked.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Security Sentinel during PR #426 review.
