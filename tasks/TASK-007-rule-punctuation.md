# Rule: Punctuation

- **ID**: TASK-007  
- **Owner**: Backend Lead  
- **Size**: M  
- **Confidence**: High  
- **Hard Dependencies**: TASK-002  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Grammar & Correctness  

## Mission Profile
- Implement punctuation fixer.  
- Remove spaces before `. , ! ? : ;`.  
- Preserve multi-punctuation clusters.  

## Deliverables
- `internal/rules/punctuation.go`.  
- Unit tests for simple and complex punctuation.  
- Golden tests T7.  
- Tricky test C4.  

## Acceptance Criteria
- ✅ `I was sitting over there ,and then BAMM !!` → `I was sitting over there, and then BAMM!!`.  
- ✅ `I waited ... and then ?!` → `I waited... and then?!`.  

## Verification Plan
- `unit`: Tests for spacing rules.  
- `integration`: Golden T7.  
- `e2e`: CLI run with punctuation fixes.  

## References
- `docs/Analysis.md` Rule Catalog (punctuation).  

## Notes for Codex Operator
- Preserve ellipsis and mixed punctuation order.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Rule: Punctuation (TASK-007)**.

### Step 1 — Analyze & Confirm
- Review punctuation fix rules.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Golden T7. Tricky C4.

### Step 3 — Generate the Code
- Implement punctuation fixer.

### Step 4 — QA & Mark Complete
- Run unit and CLI tests.
- If all pass, output: **“✅ Rule: Punctuation (TASK-007) self-verified. Please approve to mark Done.”**
