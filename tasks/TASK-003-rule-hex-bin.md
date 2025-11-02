# TASK-003 — Rule: Hex & Bin Conversion

Category: Rules  
Stage: Implementation → Testing → Refactor  
Priority: High  
Owner: tsalou23  
Created: 2025-10-30  
Last Updated: 2025-11-02  
Auditor: TBD

---

## Analyze

Implement numeric transformations for:
- (hex) → convert the previous word from hexadecimal to decimal.
- (bin) → convert the previous word from binary to decimal.

Replace the word before each marker with its correct decimal value.

Examples:
- 1E (hex) → 30  
- 10 (bin) → 2

Edge considerations:
- Ignore invalid inputs (non-hex or non-binary).
- Must handle punctuation after numbers.
- Should integrate smoothly with the pipeline.

---

## Tests

| ID | Input | Expected Output | Description |
|----|--------|----------------|--------------|
| T1 | 1E (hex) files were added | 30 files were added | Basic hexadecimal conversion |
| T2 | It has been 10 (bin) years | It has been 2 years | Basic binary conversion |
| C2 | 10 (bin) and 1A (hex) | 2 and 26 | Mixed numeric conversions |

Golden Tests: T1, T2  
Tricky Tests: C2  

Notes:
- T1 and T2 ensure correct standalone conversions.
- C2 checks mixed conversions.
- Add future tests for invalid cases like ZZ (hex) or 102 (bin).

Testing Strategy:
- Unit tests in tests/golden_test.go and tests/tricky_test.go.
- Integration tested through the pipeline (Task 002).
- Use t.Run("T1", ...) subtests to match IDs.

---

## Implement

Files:
- internal/rules/numbers.go
- tests/golden_test.go
- tests/tricky_test.go

Functions to implement:
- ApplyNumbers(text string) string  
- convertHex(word string) (string, bool)  
- convertBin(word string) (string, bool)

Implementation hints:
- Use strconv.ParseInt(word, 16, 64) for hex.
- Use strconv.ParseInt(word, 2, 64) for bin.
- Detect markers (hex) and (bin) while iterating over tokens.
- Replace the previous word and remove the marker.

---

## Acceptance

- Golden tests T1 and T2 pass.  
- Tricky test C2 passes.  
- Integrates correctly with the pipeline (Task 002).  
- Handles invalid inputs gracefully (no crash).  

---

## Refactor

To complete after tests pass:
- Extract shared logic into helper for numeric conversions.  
- Rename variables to follow Go conventions.  
- Remove debug prints.  
- Re-run go test -v ./...  

---

## Status

In Progress
