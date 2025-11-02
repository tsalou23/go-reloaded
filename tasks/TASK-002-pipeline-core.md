# TASK-002 — Pipeline Core

**Category:** Processor  
**Stage:** Implementation → Testing → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Design the baseline modular processor (pipeline).  
Sequence: tokenize → apply rules → fix articles → clean quotes → fix punctuation → join output.

## Tests
- Verify each step is called in correct order  
- Validate rule integration (e.g. hex, up, low)  

## Implement
- Files:  
  - `internal/processor/processor.go` (interface)  
  - `internal/processor/pipeline.go`  
- Function: `ProcessPipeline(text string) string`

## Acceptance
- ✅ End-to-end flow produces expected Golden results  
- ✅ All rule modules integrated  

## Dependencies
- Requires: TASK-001 (CLI)  
- Provides: Base for FSM & Hybrid

## Refactor
_(To be completed after all tests pass)_

## Status
- 🚧 In Progress  
