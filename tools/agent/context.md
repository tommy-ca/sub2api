# Agent Context: SSRF and Network Testing

This file tracks the status and patterns for SSRF protection verification and network connectivity testing.

## Deny-by-Default Policy (SSRF)

As of 2026, the application follows a strict "Deny-by-Default" policy for outbound HTTP requests:
- All private IP ranges (RFC 1918, RFC 4193) are blocked.
- All multicast ranges are blocked.
- All transition mechanism ranges (Teredo, 6to4) are blocked.
- Unspecified addresses (`0.0.0.0`, `::`) and the entire `0.0.0.0/8` range are blocked.
- Verification must happen at the **Dial level** (`net.Dialer.Control`) to prevent DNS rebinding.

## Patterns

### SSRF Testing
- Use `net.Dialer` with the `Control` function directly for unit testing SSRF blocks without needing a full HTTP lifecycle.
- See `backend/internal/pkg/httpclient/ssrf_test.go` for the canonical example.

### Network Connectivity
- Connectivity is verified using the Omakase-style script `bin/check_connectivity`.
- This script should be run inside the container to verify service discovery via container names (`postgres`, `redis`).
- Standard container healthchecks in `docker-compose.yml` also enforce this.

## Status
- [x] Implement `envutil.GetBool` resilience tests (Task 1)
- [x] Implement exhaustive SSRF block tests in `ssrf_test.go` (Task 2)
- [x] Expand `urlvalidator` ranges (0.0.0.0/8, Teredo, 6to4, AWS IMDSv2) (Task 3)
- [x] Create and verify `bin/check_connectivity` (Task 4)
