# TASK-001 — CLI Basics

**Category:** Core / CLI  
**Stage:** Implementation → Testing → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Implement a command-line interface to run the program as:
go-reloaded <input_file> <output_file> <mode>
where `<mode>` = `pipeline`, `fsm`, or `hybrid`.

Handle:
- Missing arguments → usage help  
- Invalid mode → error message  
- File I/O errors → proper exit code  

---

## Tests
| ID | Scenario | Expected Behavior |
|----|-----------|------------------|
| CLI-1 | Missing arguments | Prints usage message, exits code 1 |
| CLI-2 | Invalid mode | Prints “invalid mode” error, exits code 1 |
| CLI-3 | I/O error (bad path) | Exits with code 2 |

**Notes:**  
- Validate with integration test (`tests/cli_test.go`)  

**Testing Strategy:**  
- Use `os.Args` manipulation in test environment.  
- Run `go test ./tests/cli_test.go`.

---

## Implement
- File: `cmd/go-reloaded/main.go`
- Functions: argument parser, usage printer
- Exit codes:  
  - `1` → bad usage  
  - `2` → file I/O error  

---

## Acceptance
- ✅ CLI behaves according to spec  
- ✅ Error handling validated  

---

## Refactor
_(To be completed after all tests pass)_

---

## Status
- 🚧 In Progress  
