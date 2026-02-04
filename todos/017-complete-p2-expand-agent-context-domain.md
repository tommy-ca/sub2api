---
status: complete
priority: p2
issue_id: "017"
tags: [agent-native, documentation]
dependencies: []
---

# Problem Statement
`tools/agent/context.md` is currently "Infrastructure-Native" but lacks app-specific domain context.

# Findings
- Agents can test networking but don't understand the "Announcement" or "Group" resources.
- Agent-Native reviewer flagged this documentation gap.

# Proposed Solutions

## Option 1: Expand Domain Capability Map (Recommended)
Add sections for "Domain Resources" and "Capabilities" (primary API paths) to the agent context.
- **Pros**: Helps agents proactively manage the system.
- **Cons**: None.
- **Effort**: Small
- **Risk**: Low

# Acceptance Criteria
- [x] `tools/agent/context.md` includes primary API endpoints and domain resource descriptions.

# Work Log
## 2026-02-04 - Initial Finding
Detected by Agent-Native Reviewer.

## 2026-02-04 - Resolved
Added Domain Resources (Announcements, Groups, Accounts) and Capabilities (Gateway, Admin API) to `tools/agent/context.md`.
