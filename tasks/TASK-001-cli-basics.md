# CLI Basics

- **ID**: TASK-001
- **Owner**: Backend Lead
- **Size**: S
- **Confidence**: High
- **Hard Dependencies**: None
- **Soft Dependencies**: None
- **Related Blueprint Pillars**: Developer Experience, CLI usability

## Mission Profile
- Implement the CLI entry point for Go-Reloaded.  
- Parse `<input_file> <output_file> <mode>`.  
- Validate mode values and return appropriate exit codes.  
- Provide usage help when arguments are missing or invalid.

## Deliverables
- `/cmd/go-reloaded/main.go` with argument parsing and error handling.  
- Usage help message with modes listed.  
- Exit codes: `1` for usage errors, `2` for file I/O errors.  

## Acceptance Criteria
- ✅ Running with no args prints usage and exits with code `1`.  
- ✅ Running with invalid mode prints:  Error: invalid mode. Use one of [pipeline|fsm|hybrid]
- ✅ Running with invalid file path exits with code `2`.  
- ✅ Running with valid input/output runs without crash.  

## Verification Plan
- `unit`: Test arg parsing with valid and invalid args.  
- `integration`: CLI run with sample file → output file generated.  
- `e2e`: Shell script validating exit codes and messages.  

## References
- `docs/Analysis.md` CLI Mode Specification.  

## Notes for Codex Operator
- Keep CLI minimal; detailed processing logic comes later.  
- Ensure output messages match exactly what’s in Analysis.md.  

## PROMPT — FULL 4-STEP FLOW (execute sequentially)

You are GPT-Codex executing **CLI Basics (TASK-001)**.

### Step 1 — Analyze & Confirm
- Review CLI requirements in `docs/Analysis.md`.
- Summarize expected args, modes, and error handling.
- WAIT for confirmation.

### Step 2 — Generate the Tests
- Write tests for no args, invalid mode, invalid file, valid run.
- Capture exit codes and stdout/stderr.

### Step 3 — Generate the Code
- Implement `main.go` parsing logic, usage help, error handling.

### Step 4 — QA & Mark Complete
- Run CLI manually and automated tests.
- If all pass, output: **“✅ CLI Basics (TASK-001) self-verified. Please approve to mark Done.”**

