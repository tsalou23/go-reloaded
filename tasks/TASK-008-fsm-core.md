# FSM Core

- **ID**: TASK-008  
- **Owner**: Backend Lead  
- **Size**: M  
- **Confidence**: Medium  
- **Hard Dependencies**: TASK-001  
- **Soft Dependencies**: TASK-002  
- **Related Blueprint Pillars**: Architecture Foundations  

## Mission Profile
- Implement Finite State Machine processor for parsing.  
- Handle tokens, markers, quotes, and punctuation sequences.  

## Deliverables
- `internal/processor/fsm.go`.  
- FSM states: Normal, MarkerFound, InQuotes, InPunctuation.  
- CLI runs in `fsm` mode with identity output.  

## Acceptance Criteria
- ✅ CLI runs in fsm mode without crash.  
- ✅ Input text returned unchanged.  
- ✅ FSM states initialized correctly.  

## Verification Plan
- `unit`: FSM transitions tested.  
- `integration`: CLI run with FSM mode identity transform.  

## References
- `docs/ARCHITECTURE.md` FSM section.  

## Notes for Codex Operator
- Start with skeleton FSM, add rules later.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **FSM Core (TASK-008)**.

### Step 1 — Analyze & Confirm
- Review FSM design.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Unit tests for state transitions.

### Step 3 — Generate the Code
- Implement FSM skeleton.

### Step 4 — QA & Mark Complete
- Run FSM in CLI, confirm identity.
- If all pass, output: **“✅ FSM Core (TASK-008) self-verified. Please approve to mark Done.”**
