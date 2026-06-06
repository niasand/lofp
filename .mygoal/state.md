# Goal: Complete i18n integration (3 remaining tasks)

## Status: completed
## Created: 2026-06-06 23:00
## Updated: 2026-06-06 23:45

## Objective
Complete the 3 remaining i18n tasks for LoFP:
1. [P0] Add missing translation entries for weather/time/monsters/treasure to messages_zh.go
2. [P1] Integrate zh.json into frontend React components
3. [P2] Add LOCALE=zh to fly.toml for production

## Verification
- `cd engine && go build ./...` passes
- `cd frontend && npx tsc --noEmit` passes
- All new strings have Chinese translations in messages_zh.go
- Frontend components use zh.json translations
- fly.toml contains LOCALE=zh

## Constraints
- Don't break existing functionality
- Keep format placeholders (%s, %d, %%, %N, %T) intact
- Match existing code style

## Boundaries
- engine/internal/engine/ (weather.go, gametime.go, monsters.go, treasure.go, player.go)
- frontend/src/ (components that have English strings)
- fly.toml
- engine/i18n/messages_zh.go

## Iteration Policy
Start with P0 (add translations), then P1 (frontend integration), then P2 (fly.toml). Each step must compile before moving on.

## Blocked Stop Condition
If a file can't be modified without breaking the build, report the blocker and the specific error.

---

## Evidence Ledger

| # | Claim | Evidence | Status |
|---|-------|----------|--------|
| 1 | i18n infrastructure exists | engine/i18n/i18n.go, messages_zh.go, frontend/src/i18n/zh.json | confirmed |
| 2 | 247 strings already wrapped | 7 commits from subagents | confirmed |
| 3 | weather/time/monsters/treasure have no translations | subagent report: "none of the string categories in these five files" | confirmed |
| 4 | Frontend zh.json exists but not integrated | zh.json created but no components import it | confirmed |
| 5 | fly.toml has no LOCALE setting | fly.toml only has LOFP_CONFIG, LOFP_STATIC_DIR, MONGODB_DATABASE, GOMEMLIMIT | confirmed |
| 6 | 113 new translations added | P0 agent: weather(34), time(14), races(9), directions(11), pronouns(10), monsters(8), treasure(8), movement(4) | confirmed |
| 7 | 5 source files wrapped with i18n.T() | weather.go, gametime.go, monsters.go, treasure.go, player.go — go build passes | confirmed |
| 8 | Frontend 7 components translated | AdminPanel, MainMenu, CaptureModal, CaptureViewer, ResetPassword, VerifyEmail, APIDocs — tsc passes | confirmed |
| 9 | fly.toml has LOCALE=zh + VITE_LOCALE=zh | P2 agent committed | confirmed |

## Iteration Log

| # | Time | Action | Result | Next |
|---|------|--------|--------|------|
| 1 | 23:00 | Created goal state | Goal initialized | Start P0: add missing translations |
| 2 | 23:45 | P0+P1+P2 all completed | All 3 tasks done, builds pass | Goal complete |
