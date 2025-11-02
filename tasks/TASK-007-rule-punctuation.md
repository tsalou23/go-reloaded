# TASK-007 — Rule: Punctuation

**Category:** Rules  
**Stage:** Implementation → Testing → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Ensure punctuation `. , ! ? : ;` has no space before and one space after,  
except clusters like `...`, `?!`, `!!`.

---

## Tests
| ID | Input | Expected Output |
|----|--------|----------------|
| T7 | `I was sitting over there ,and then BAMM !!` | `I was sitting over there, and then BAMM!!` |
| C4 | `I waited ... and then ?!` | `I waited... and then?!` |

**Golden Tests:** T7  
**Tricky Tests:** C4  

**Notes:**  
- Check spacing near quotes too.  

---

## Implement
- File: `internal/rules/punctuation.go`
- Function: `FixPunctuation(text string) string`
- Use regex: `\s*([.,!?;:]+)\s*`

---

## Acceptance
- ✅ Golden + Tricky tests pass  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- 🚧 In Progress  
