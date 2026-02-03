---
status: complete
priority: p3
issue_id: "005"
tags: [security, networking, go]
dependencies: []
---

# Problem Statement
The current SSRF protection resolves IPs before the request. This leaves a theoretical TOCTOU window where DNS could change between the check and the actual connection.

# Findings
- `backend/internal/pkg/httpclient/pool.go` (implied context) performs validation at the transport level.
- Hardening at the `DialContext` level would be more robust.

# Proposed Solutions

## Option 1: Custom Dialer Validation
Implement a custom `DialContext` that validates the IP address immediately after it is resolved by the dialer but before the TCP handshake is initiated.
- **Pros**: Eliminates TOCTOU risk.
- **Cons**: More complex implementation.
- **Effort**: Medium
- **Risk**: Low

# Acceptance Criteria
- [x] DialContext-level validation implemented.
- [x] Tests covering DNS rebinding scenarios.

# Work Log
## 2026-02-03 - Initial Finding
Detected by security-sentinel agent.

## 2026-02-03 - Resolved
Implemented robust Dial-level IP validation in `httpclient/pool.go`.
