---
status: complete
priority: p1
issue_id: "009"
tags: [security, setup]
dependencies: []
---

# Problem Statement
If the application fails to write the `.initial_admin_password` file, it currently logs the password to stdout as a fallback. This is a security risk.

# Findings
- `backend/internal/setup/setup.go` has a fallback that uses `fmt.Printf` if `os.WriteFile` fails.
- This results in sensitive credentials persisting in container logs.

# Proposed Solutions

## Option 1: Fail Fast (Recommended)
Replace the fallback with `log.Fatalf`. If the system cannot securely store the initial password, it should refuse to start.
- **Pros**: Prevents credential leakage.
- **Cons**: Might cause startup failure in misconfigured environments (e.g., read-only data volumes).
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [ ] App calls `log.Fatalf` if `.initial_admin_password` cannot be written.
- [ ] No credential leakage in logs.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Kieran and Security Sentinel during PR #426 review.
