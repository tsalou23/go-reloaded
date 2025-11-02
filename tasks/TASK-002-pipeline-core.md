# TASK-002 — Pipeline Core

**Category:** Processor  
**Stage:** Implementation → Integration → Refactor  
**Priority:** High  
**Owner:** tsalou23  
**Created:** 2025-10-30  
**Last Updated:** 2025-11-02  
**Auditor:** _TBD_

---

## Analyze
Implement the modular **pipeline** orchestrator that applies all rule processors in order:
1. NumericFixer (hex, bin)
2. CaseFixer (up, low, cap)
3. ArticleFixer (a → an) 
4. QuoteCleaner
5. PunctuationFixer

---

## Tests
| ID | Scenario | Expected Behavior |
|----|-----------|------------------|
| PIPE-1 | All rule functions called sequentially | Output matches combination of rule results |
| PIPE-2 | Empty input | Returns empty string |
| PIPE-3 | Mixed rules (integration) | Matches paragraph expected output |

**Notes:**  
- Use mock functions for rules before actual implementation.  
- Integration verified after Tasks 003-007 complete.

**Testing Strategy:**  
- Implement tests in `tests/paragraph_test.go`.

---

## Implement
- File: `internal/processor/pipeline.go`
- Interface: `Processor`
- Function:  
  ```go
  func (p *Pipeline) Process(text string) string {
      text = rules.ApplyNumbers(text)
      text = rules.ApplyCase(text)
      text = rules.FixArticles(text)
      text = rules.CleanQuotes(text)
      text = rules.FixPunctuation(text)
      return text
  }
Acceptance

✅ Executes all rule stages in order

✅ Ready for integration with FSM & Hybrid

Refactor

(To be completed after all tests pass)

Status

🚧 In Progress