---
status: complete
priority: p1
issue_id: "007"
tags: [security, config]
dependencies: []
---

# Problem Statement
The default value for `SECURITY_URL_ALLOWLIST_ALLOW_INSECURE_HTTP` is inconsistent between the code and the `.env.example` file. 

# Findings
- `backend/internal/config/config.go` sets the default to `false`.
- `deploy/.env.example` still has it set to `true`.
- This creates confusion and might lead to insecure deployments by default if users copy the example.

# Proposed Solutions

## Option 1: Update .env.example (Recommended)
Change the value in `deploy/.env.example` to `false` to match the code's secure default.
- **Pros**: Consistency, secure by default.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [ ] `deploy/.env.example` has `SECURITY_URL_ALLOWLIST_ALLOW_INSECURE_HTTP=false`.

# Work Log
## 2026-02-03 - Initial Finding
Detected by Kieran during PR #426 review.
