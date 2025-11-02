# TASK-010 — Golden & Tricky Tests

**Category:** Testing  
**Stage:** Implementation → Execution → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** TBD

---

## Analyze

Collect and execute all test cases for the Go-Reloaded project:
- Golden tests (T1–T10)
- Tricky tests (C1–C5)
- Full paragraph integration test
- CLI behavior tests

The goal is to ensure that all transformation rules (hex/bin, case, articles, quotes, punctuation) work correctly and consistently across the three processor modes: **Pipeline**, **FSM**, and **Hybrid**.

---

## Tests

### Golden Tests (T1–T10)
| ID | Input | Expected Output |
|----|--------|----------------|
| T1 | 1E (hex) files were added | 30 files were added |
| T2 | It has been 10 (bin) years | It has been 2 years |
| T3 | Ready, set, go (up) ! | Ready, set, GO! |
| T4 | I should stop SHOUTING (low) | I should stop shouting |
| T5 | Welcome to the brooklyn bridge (cap) | Welcome to the Brooklyn Bridge |
| T6 | This is so exciting (up, 2) | This is SO EXCITING |
| T7 | I was sitting over there ,and then BAMM !! | I was sitting over there, and then BAMM!! |
| T8 | There it was. A amazing rock! | There it was. An amazing rock! |
| T9 | I am exactly how they describe me: ' awesome ' | I am exactly how they describe me: 'awesome' |
| T10 | As Elton John said: ' I am the most well-known homosexual in the world ' | As Elton John said: 'I am the most well-known homosexual in the world' |

### Tricky Tests (C1–C5)
| ID | Input | Expected Output |
|----|--------|----------------|
| C1 | a honest man | an honest man |
| C2 | 10 (bin) and 1A (hex) | 2 and 26 |
| C3 | HELLO (low, 2) WORLD | hello world WORLD |
| C4 | I waited ... and then ?! | I waited... and then?! |
| C5 | He said ' hello there ' | He said 'hello there' |

### Integration Test
- Based on the **Large Paragraph Example** in `docs/Analysis.md` section 7.  
- Ensures all rules interact properly in long-form text.

### CLI Tests
| ID | Scenario | Expected Result |
|----|-----------|----------------|
| CLI-1 | Missing arguments | Shows usage message, exits code 1 |
| CLI-2 | Invalid mode | Shows error message, exits code 1 |
| CLI-3 | File I/O error | Exits with code 2 |

**Testing Strategy**
- `tests/golden_test.go` → T1–T10  
- `tests/tricky_test.go` → C1–C5  
- `tests/paragraph_test.go` → full integration test  
- `tests/cli_test.go` → CLI behavior  
- Each test uses `t.Run("T#", ...)` or `t.Run("C#", ...)` for traceability.  
- Use `go test -v ./tests/...` to run all tests.

---

## Implement

**Files**
- `tests/golden_test.go`  
- `tests/tricky_test.go`  
- `tests/paragraph_test.go`  
- `tests/cli_test.go`  

**Example Test Function**

```go
func TestGolden_T1(t *testing.T) {
    input := "1E (hex) files were added"
    expected := "30 files were added"
    output := ProcessPipeline(input)
    if output != expected {
        t.Errorf("T1 failed: got %q, want %q", output, expected)
    }
}
```

**Implementation Notes**
- Each Markdown test ID maps directly to a `t.Run()` block in Go.
- All three processing modes (pipeline, fsm, hybrid) must produce identical output.
- Ensure tests are deterministic (no randomness or file dependency).

---

## Acceptance

- All Golden tests (T1–T10) pass.  
- All Tricky tests (C1–C5) pass.  
- Paragraph integration test passes across all modes.  
- CLI tests return correct error messages and exit codes.  
- Test coverage ≥ 90% for `internal/` directory.

---

## Refactor

After all tests pass:
- Merge duplicate test logic between golden and tricky test files.  
- Extract shared helpers for input/output assertions.  
- Add benchmark tests for performance comparison.  
- Run `go test -cover ./...` and document coverage metrics.

---

## Status

In Progress
