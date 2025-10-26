# Rule: Case Formatting

- **ID**: TASK-004  
- **Owner**: Backend Lead  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Rule Engine  

## Mission Profile
- Implement word case rules: `(up)`, `(low)`, `(cap)`.  
- Extend with `(up, n)`, `(low, n)`, `(cap, n)` for multi-word formatting.  

## Deliverables
- `internal/rules/cases.go` with case functions.  
- Unit tests for single and multi-word rules.  
- Golden tests T3, T4, T5, T6.  
- Tricky test C3.  

## Acceptance Criteria
- ✅ `go (up)` → `GO`.  
- ✅ `LOUD (low)` → `loud`.  
- ✅ `bridge (cap)` → `Bridge`.  
- ✅ `so exciting (up, 2)` → `SO EXCITING`.  
- ✅ `HELLO (low, 2) WORLD` → `hello world WORLD`.  

## Verification Plan
- `unit`: Test helper functions for casing.  
- `integration`: Golden T3–T6.  
- `e2e`: CLI run with multi-word case rules.  

## References
- `docs/Analysis.md` Rule Catalog (case).  

## Notes for Codex Operator
- Pay attention to multi-word counts and edge cases (punctuation adjacency).  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Rule: Case Formatting (TASK-004)**.

### Step 1 — Analyze & Confirm
- Review rules for single vs multi-word casing.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Golden T3–T6. Tricky C3.

### Step 3 — Generate the Code
- Implement case functions in cases.go.

### Step 4 — QA & Mark Complete
- Run unit and CLI tests.
- If all pass, output: **“✅ Rule: Case Formatting (TASK-004) self-verified. Please approve to mark Done.”**
