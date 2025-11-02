# TASK-008 — FSM Core

**Category:** Processor  
**Stage:** Implementation → Testing → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Create a Finite State Machine (FSM) for parsing input and detecting markers `(hex)`, `(up)`, etc.

---

## Tests
| ID | Scenario | Expected Behavior |
|----|-----------|------------------|
| FSM-1 | Correctly transitions NORMAL → MARKER → APPLY | Tokens recognized properly |
| FSM-2 | Handles quotes `' ... '` as atomic | Tokens inside quotes untouched |

**Notes:**  
- Compare FSM output with tokenizer baseline.  

---

## Implement
- File: `internal/processor/fsm.go`
- Structs: `State`, `Transition`
- FSM states: NORMAL, MARKER_FOUND, APPLY_RULE

---

## Acceptance
- ✅ FSM parses input correctly  
- ✅ Matches pipeline tokenizer output  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- 🚧 In Progress  
