# Rule: Quotes Handling

- **ID**: TASK-006  
- **Owner**: Backend Lead  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Grammar & Correctness  

## Mission Profile
- Implement removal of unnecessary spaces inside `' '`.  

## Deliverables
- `internal/rules/quotes.go` with quote cleaner.  
- Unit tests.  
- Golden tests T9, T10.  
- Tricky test C5.  

## Acceptance Criteria
- ✅ `I am ' happy '` → `I am 'happy'`.  
- ✅ `He said ' hello there '` → `He said 'hello there'`.  

## Verification Plan
- `unit`: Test cleaner on single and multi-word cases.  
- `integration`: Golden T9–T10.  
- `e2e`: CLI run with quotes in text.  

## References
- `docs/Analysis.md` Rule Catalog (quotes).  

## Notes for Codex Operator
- Preserve punctuation outside quotes.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Rule: Quotes Handling (TASK-006)**.

### Step 1 — Analyze & Confirm
- Review quote handling requirements.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Golden T9, T10. Tricky C5.

### Step 3 — Generate the Code
- Implement quote cleaner.

### Step 4 — QA & Mark Complete
- Run unit and CLI tests.
- If all pass, output: **“✅ Rule: Quotes Handling (TASK-006) self-verified. Please approve to mark Done.”**
