# TASK-004 — Rule: Case Formatting

**Category:** Rules  
**Stage:** Implementation → Testing → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Apply `(up)`, `(low)`, `(cap)` and counted forms `(up, n)`, `(low, n)`, `(cap, n)`.

---

## Tests
| ID | Input | Expected Output |
|----|--------|----------------|
| T3 | `Ready, set, go (up) !` | `Ready, set, GO!` |
| T4 | `I should stop SHOUTING (low)` | `I should stop shouting` |
| T5 | `Welcome to the brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |
| T6 | `This is so exciting (up, 2)` | `This is SO EXCITING` |
| C3 | `HELLO (low, 2) WORLD` | `hello world WORLD` |

**Golden Tests:** T3–T6  
**Tricky Tests:** C3  

**Notes:**  
- Verify counted forms and punctuation handling.  

---

## Implement
- File: `internal/rules/cases.go`
- Helper: `applyCase(word, mode)`
- Exported: `ApplyCase(text string) string`

---

## Acceptance
- ✅ Golden tests T3–T6 pass  
- ✅ Tricky test C3 pass  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- ✅ Done  
