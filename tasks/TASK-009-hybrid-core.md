# Hybrid Core

- **ID**: TASK-009  
- **Owner**: Backend Lead  
- **Size**: M  
- **Confidence**: Medium  
- **Hard Dependencies**: TASK-002, TASK-008  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Architecture Foundations  

## Mission Profile
- Implement hybrid processor.  
- Use FSM for tokenization + pipeline for rules.  

## Deliverables
- `internal/processor/hybrid.go`.  
- CLI runs in `hybrid` mode with identity transform.  

## Acceptance Criteria
- ✅ CLI runs in hybrid mode without crash.  
- ✅ Input text returned unchanged.  
- ✅ FSM tokenizer integrated with pipeline rules.  

## Verification Plan
- `unit`: Test FSM + pipeline integration.  
- `integration`: CLI run hybrid mode identity.  

## References
- `docs/ARCHITECTURE.md` Hybrid section.  

## Notes for Codex Operator
- Modular integration between FSM and pipeline.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Hybrid Core (TASK-009)**.

### Step 1 — Analyze & Confirm
- Review hybrid integration needs.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Test tokenizer → pipeline passthrough.

### Step 3 — Generate the Code
- Implement hybrid processor.

### Step 4 — QA & Mark Complete
- Run hybrid CLI and confirm identity.
- If all pass, output: **“✅ Hybrid Core (TASK-009) self-verified. Please approve to mark Done.”**
