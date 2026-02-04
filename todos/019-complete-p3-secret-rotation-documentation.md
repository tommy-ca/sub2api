---
status: complete
priority: p3
issue_id: "019"
tags: [documentation, security]
dependencies: []
---

# Problem Statement
With auto-generated secrets, users need a clear guide on how to rotate them safely.

# Findings
- Architecture Strategist recommended providing a rotation guide for `JWT_SECRET` and `TOTP_ENCRYPTION_KEY`.

# Proposed Solutions

## Option 1: Security Maintenance Guide
Create a new section in `deploy/README.md` or a new `SECURITY.md` detailing rotation impacts.
- **Pros**: Better operational security.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] Documentation exists for secret rotation.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Architecture Strategist.

## 2026-02-04 - Documentation Created
- Created `SECURITY.md` with detailed rotation guides for `JWT_SECRET` and `TOTP_ENCRYPTION_KEY`.
- Linked to `SECURITY.md` from `deploy/README.md`.
