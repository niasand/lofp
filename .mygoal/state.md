# Goal: Find all untranslated (hardcoded English) strings in frontend

## Status: completed
## Created: 2026-06-07 22:30
## Updated: 2026-06-07 22:30

## Objective
Produce a comprehensive list of all hardcoded English strings in the frontend React components that are NOT going through the i18n translation system (useTranslation / t() function). Every string that a Chinese-speaking user would see should be translatable.

## Verification
Every component .tsx file scanned, each hardcoded string identified with file:line, and the list is complete.

## Constraints
- Don't modify any code, just audit and list
- Focus on user-visible strings (not CSS classes, variable names, technical identifiers)

## Boundaries
frontend/src/components/*.tsx, frontend/src/App.tsx

## Iteration Policy
Fan out agents to scan different component groups in parallel, then consolidate findings.

## Blocked Stop Condition
If i18n setup is unclear, report blocker.

---

## Evidence Ledger

| # | Claim | Evidence | Status |
|---|-------|----------|--------|

## Iteration Log

| # | Time | Action | Result | Next |
|---|------|--------|--------|------|
| 1 | 22:30 | Fan out agents to scan all components | Pending | Consolidate results |
