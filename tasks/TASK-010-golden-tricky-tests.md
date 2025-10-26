# Golden & Tricky Tests

- **ID**: TASK-010  
- **Owner**: QA Lead  
- **Size**: L  
- **Confidence**: High  
- **Hard Dependencies**: TASK-003 → TASK-007  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Quality & Verification  

## Mission Profile
- Implement Golden and Tricky test sets.  
- Ensure all rules pass before acceptance.  

## Deliverables
- `tests/golden_test.go`.  
- `tests/tricky_test.go`.  
- `tests/paragraph_test.go`.  
- All T1–T10, C1–C5, and Large Paragraph.  

## Acceptance Criteria
- ✅ All Golden tests pass.  
- ✅ All Tricky tests pass.  
- ✅ Large Paragraph matches expected output.  

## Verification Plan
- `unit`: Test individual rules.  
- `integration`: Golden & Tricky tests.  
- `e2e`: Paragraph test with full pipeline.  

## References
- `docs/Analysis.md` Golden Test Set, Tricky Cases, Large Paragraph.  

## Notes for Codex Operator
- Lock in expected outputs; use as regression suite.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Golden & Tricky Tests (TASK-010)**.

### Step 1 — Analyze & Confirm
- Review Golden, Tricky, Paragraph expected outputs.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Implement T1–T10, C1–C5, Paragraph.

### Step 3 — Generate the Code
- Wire tests to processors.

### Step 4 — QA & Mark Complete
- Run `go test ./tests/...`.
- If all pass, output: **“✅ Golden & Tricky Tests (TASK-010) self-verified. Please approve to mark Done.”**
