# Rule: Hex & Bin

- **ID**: TASK-003  
- **Owner**: Backend Lead  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Rule Engine  

## Mission Profile
- Implement numeric conversion rules: `(hex)` and `(bin)`.  
- Convert previous word into decimal.  

## Deliverables
- `internal/rules/numbers.go` with conversion logic.  
- Unit tests for `(hex)` and `(bin)`.  
- Golden tests T1, T2.  
- Tricky test C2.  

## Acceptance Criteria
- ✅ `1E (hex)` → `30`.  
- ✅ `10 (bin)` → `2`.  
- ✅ `10 (bin) and 1A (hex)` → `2 and 26`.  

## Verification Plan
- `unit`: Numbers.go conversion functions.  
- `integration`: Golden tests T1, T2.  
- `e2e`: CLI run with sample file containing hex/bin markers.  

## References
- `docs/Analysis.md` Rule Catalog (hex/bin).  

## Notes for Codex Operator
- Validate error handling for invalid numbers.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Rule: Hex & Bin (TASK-003)**.

### Step 1 — Analyze & Confirm
- Review hex/bin conversion requirements.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Golden T1, T2. Tricky C2.

### Step 3 — Generate the Code
- Implement numbers.go conversion functions.

### Step 4 — QA & Mark Complete
- Run unit and CLI tests.
- If all pass, output: **“✅ Rule: Hex & Bin (TASK-003) self-verified. Please approve to mark Done.”**
