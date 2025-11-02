# TASK-005 — Rule: Articles (a → an)

**Category:** Rules  
**Stage:** Implementation → Testing → Refactor  
**Priority:** Medium  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Replace “a” with “an” when next word starts with a vowel or `h`.

---

## Tests
| ID | Input | Expected Output |
|----|--------|----------------|
| T8 | `There it was. A amazing rock!` | `There it was. An amazing rock!` |
| C1 | `a honest man` | `an honest man` |

**Golden Tests:** T8  
**Tricky Tests:** C1  

**Notes:**  
- Add case sensitivity check (`A honest mistake`).  

---

## Implement
- File: `internal/rules/articles.go`
- Function: `FixArticles(text string) string`
- Regex or manual string scan for word boundaries.

---

## Acceptance
- ✅ Golden + Tricky tests pass  
- ✅ Handles vowels and `h` prefix  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- ✅ Done  
