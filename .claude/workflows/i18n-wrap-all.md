# i18n Wrap All Hardcoded Strings

Wrap all player-facing hardcoded English strings in `i18n.T()` across the engine codebase.

## Workflow

1. 12 parallel agents each handle one file or file group
2. Each agent wraps strings in i18n.T() and returns new translation keys
3. Final agent consolidates all new keys into messages_zh.go
