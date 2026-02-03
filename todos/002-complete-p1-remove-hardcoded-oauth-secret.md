---
status: complete
priority: p1
issue_id: "002"
tags: [security, oauth, go]
dependencies: []
---

# Problem Statement
The `antigravity` OAuth implementation contains a hardcoded `ClientSecret`. This secret should be treated as a sensitive credential and managed via environment variables or the configuration system.

# Findings
- `backend/internal/pkg/antigravity/oauth.go` has a hardcoded string for the client secret.
- This was flagged by the security-sentinel agent as a high-severity risk outside the direct scope of PR #426 but relevant to the overall security posture.

# Proposed Solutions

## Option 1: Move to Environment Variables
Define a new environment variable `ANTIGRAVITY_OAUTH_CLIENT_SECRET` and update the config loader to inject it into the service.
- **Pros**: Standardized secret management.
- **Cons**: Requires users to update their `.env` files.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] Hardcoded secret removed from source code.
- [x] Secret is loaded from environment/config.
- [x] `.env.example` updated with the new variable.

# Work Log
## 2026-02-03 - Initial Finding
Detected during PR #426 review as pre-existing technical debt.

## 2026-02-03 - Resolved
Removed hardcoded secrets from `config.go`, `constants.go`, and `oauth.go`. Updated documentation.
