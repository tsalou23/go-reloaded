# TASK-006 — Rule: Quotes

**Category:** Rules  
**Stage:** Implementation → Testing → Refactor  
**Priority:** Medium  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Remove unnecessary spaces inside `'...'` and preserve content.  

---

## Tests
| ID | Input | Expected Output |
|----|--------|----------------|
| T9 | `I am exactly how they describe me: ' awesome '` | `I am exactly how they describe me: 'awesome'` |
| T10 | `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` |
| C5 | `He said ' hello there '` | `He said 'hello there'` |

**Golden Tests:** T9, T10  
**Tricky Tests:** C5  

**Notes:**  
- Handle multi-word quotes.  
- Consider punctuation after closing quote.

---

## Implement
- File: `internal/rules/quotes.go`
- Function: `CleanQuotes(text string) string`

---

## Acceptance
- ✅ Golden & Tricky tests pass  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- 🚧 In Progress  
