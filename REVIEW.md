# Code Review Guidelines for LoFP

This file defines review priorities and rules for the Legends of Future Past codebase. It is read by Claude Code's review agents at highest priority.

## Severity Definitions

| Level | Meaning | Action |
|-------|---------|--------|
| 🔴 **Important** | Bug that could corrupt game state, break combat, or cause data loss | Must fix before merge |
| 🟡 **Nit** | Code quality, minor style, small optimizations | Fix if convenient, don't block merge |
| 🟣 **Pre-existing** | Bug in existing code, not introduced by this PR | Note for awareness, don't block |

## Critical Review Areas

### 1. Game Logic Correctness (🔴 Important)
- Combat calculations (ToHit, damage rolls, crits) must match original `shirla.cap` session format
- Spell/psionic effects must reference `original/GMSCRIPT.DOC` for correct behavior
- Crafting recipes must match `original/GM Pages/MANUAL.DOC` specifications
- XP/level calculations must be consistent with documented tables

### 2. Concurrency & State Sync (🔴 Important)
- WebSocket handlers must not share mutable state without synchronization
- Room state changes MUST call `notifyRoomChange()` or publish through hub
- Monster spawning/combat must handle concurrent player interactions
- Presence heartbeats must not cause race conditions

### 3. Data Integrity (🔴 Important)
- MongoDB operations must handle connection failures gracefully
- Player state saves must be atomic (don't save partial updates)
- Character deletion must be soft-delete, never hard-delete without confirmation
- Item/room state mutations must be reversible or logged

### 4. Security (🔴 Important)
- All user input must be sanitized before display (HTML injection)
- Rate limiting must be enforced on all public endpoints
- JWT tokens must be validated on every authenticated request
- Admin endpoints must verify admin role
- No secrets in code or commits (check .env patterns)

### 5. Protocol Compliance (🟡 Nit)
- GMCP messages must match documented package format
- MXP secure tags only in `\033[1z` line mode
- MCCP2: ALL output through compressor when active (never raw `conn.Write()`)
- Password echo suppression must bracket password prompts

### 6. Performance (🟡 Nit)
- Avoid blocking operations in WebSocket goroutines
- Room/item lookups should use indexed queries
- Monster AI loops must have exit conditions
- Frontend re-renders should be minimized (React.memo, useMemo)

## Skip Rules

- **Skip `original/` directory** — read-only reference scripts, not code
- **Skip `.env` files** — secrets, never commit
- **Skip generated files** — `frontend/dist/`, `engine/vendor/`

## Nit Budget

- Maximum 5 Nit-level comments per review
- Prioritize actionable Nits over stylistic preferences

## Auto-Flag Patterns

Flag as 🔴 Important if found:
- `conn.Write()` after MCCP2 negotiation (should use `t.write()`)
- Missing `notifyRoomChange()` after room state mutation
- Hardcoded MongoDB collection names (use constants)
- Unchecked error returns from database operations
- `go func()` without context cancellation

Flag as 🟡 Nit:
- Functions > 100 lines
- Duplicated logic across handlers
- Missing godoc comments on exported functions
- Inconsistent error message formatting

## Reference Documents

When reviewing game mechanics changes, verify against:
1. `original/GM Pages/MANUAL.DOC` — combat, stats, items, monsters
2. `original/GMSCRIPT.DOC` — script language, room/item/monster definitions
3. `original/legends/shirla.cap` — actual gameplay output format
4. `original/legends/LEGENDS.DOC` — player-facing mechanics
