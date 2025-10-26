# Rule: Article Fix

- **ID**: TASK-005  
- **Owner**: Backend Lead  
- **Size**: S  
- **Confidence**: Medium  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Grammar & Correctness  

## Mission Profile
- Implement article correction: replace `a` with `an` before vowels and `h` words.  

## Deliverables
- `internal/rules/articles.go` with article fixer.  
- Unit tests for vowel + h cases.  
- Golden test T8.  
- Tricky test C1.  

## Acceptance Criteria
- ✅ `a amazing rock` → `an amazing rock`.  
- ✅ `a honest man` → `an honest man`.  

## Verification Plan
- `unit`: Article fixer tests.  
- `integration`: Golden T8.  
- `e2e`: CLI run with text containing article errors.  

## References
- `docs/Analysis.md` Rule Catalog (article fix).  

## Notes for Codex Operator
- Ensure case sensitivity is handled (e.g., “A apple” → “An apple”).  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Rule: Article Fix (TASK-005)**.

### Step 1 — Analyze & Confirm
- Review vowel and “h” logic.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Golden T8. Tricky C1.

### Step 3 — Generate the Code
- Implement article fixer.

### Step 4 — QA & Mark Complete
- Run unit and CLI tests.
- If all pass, output: **“✅ Rule: Article Fix (TASK-005) self-verified. Please approve to mark Done.”**
