# Pipeline Core

- **ID**: TASK-002  
- **Owner**: Backend Lead  
- **Size**: S  
- **Confidence**: High  
- **Hard Dependencies**: TASK-001  
- **Soft Dependencies**: None  
- **Related Blueprint Pillars**: Architecture Foundations  

## Mission Profile
- Implement the baseline **Pipeline processor**.  
- Sequential stages with tokenizer and joiner.  
- At this stage, rules are not applied; identity transform only.  

## Deliverables
- `internal/processor/processor.go` with Processor interface.  
- `internal/processor/pipeline.go` with `PipelineProcessor`.  
- CLI runs in `pipeline` mode successfully.  

## Acceptance Criteria
- ✅ CLI runs in pipeline mode without crashing.  
- ✅ Input text is output unchanged.  
- ✅ Unit test confirms identity transform.  

## Verification Plan
- `unit`: Pipeline returns same text given input.  
- `integration`: CLI run in pipeline mode → unchanged text.  

## References
- `docs/ARCHITECTURE.md`: Pipeline overview.  

## Notes for Codex Operator
- Keep design modular for later stages.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **Pipeline Core (TASK-002)**.

### Step 1 — Analyze & Confirm
- Review Processor interface needs.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Write identity tests (input = output).

### Step 3 — Generate the Code
- Implement Processor interface and PipelineProcessor skeleton.

### Step 4 — QA & Mark Complete
- Run CLI in pipeline mode and validate identity.
- If all pass, output: **“✅ Pipeline Core (TASK-002) self-verified. Please approve to mark Done.”**
