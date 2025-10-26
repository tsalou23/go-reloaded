# 🧠 Go-Reloaded — Analysis Document

## 📑 Table of Contents

1. [Problem Description](#1-problem-description)
2. [Rule Catalog](#2-rule-catalog)
3. [Architecture Comparison](#3-architecture-comparison)
    3.1 [Pipeline](#31-pipeline)
    3.2 [Finite State Machine (FSM)](#32-finite-state-machine-fsm)
    3.3 [Hybrid (Pipeline + FSM)](#33-hybrid-pipeline--fsm)
4. [Rule Execution Order](#4-rule-execution-order-pipeline-flow)
5. [Golden Test Set](#5-golden-test-set)
6. [Tricky Test Cases](#6-tricky-test-cases)
7. [Large Paragraph Example](#7-large-paragraph-example)
8. [CLI Mode Specification](#8-cli-mode-specification)
9. [Collaboration Model (Operator, Agent, Auditor)](#9-collaboration-model)
10. [Conclusion](#10-conclusion)
11. [Project File Structure](#11-project-file-structure)
12. [Future Extensions](#12-future-extensions)

---

## 1. Problem Description

**Go-Reloaded** is a command-line tool written in Go that processes an input text file, applies a set of transformation rules, and writes the corrected text to an output file.

The tool must correctly handle:

* **Number conversions**: Hexadecimal → decimal, Binary → decimal.
* **Word formatting**: Uppercase, lowercase, capitalization.
* **Punctuation corrections**: Proper spacing, handling multi-punctuation sequences.
* **Grammar fixes**: Article correction (“a” → “an”).
* **Quotes handling**: Remove unnecessary spaces inside `' '`.
* **Multi-word markers**: `(up, n)`, `(low, n)`, `(cap, n)`.

Development follows **TDD** with incremental tasks. Every feature must have unit tests and must pass Golden & Tricky test cases before being accepted.

---

## 2. Rule Catalog

| Rule                  | Description                                      | Example                                      |
| --------------------- | ------------------------------------------------ | -------------------------------------------- |
| **(hex)**             | Convert previous word from hexadecimal → decimal | `1E (hex)` → `30`                            |
| **(bin)**             | Convert previous word from binary → decimal      | `10 (bin)` → `2`                             |
| **(up)**              | Uppercase previous word                          | `go (up)` → `GO`                             |
| **(low)**             | Lowercase previous word                          | `LOUD (low)` → `loud`                        |
| **(cap)**             | Capitalize previous word                         | `bridge (cap)` → `Bridge`                    |
| **(up, n)**           | Uppercase previous *n* words                     | `so exciting (up, 2)` → `SO EXCITING`        |
| **(low, n)**          | Lowercase previous *n* words                     | `HELLO WORLD (low, 2)` → `hello world`       |
| **(cap, n)**          | Capitalize previous *n* words                    | `the big bridge (cap, 3)` → `The Big Bridge` |
| **Punctuation**       | Fix spaces before `. , ! ? : ;`                  | `Hi , world !` → `Hi, world!`                |
| **Multi-punctuation** | Preserve sequences                               | `Wait ... What ?!` → `Wait... What?!`        |
| **Quotes `' '`**      | Remove spaces inside quotes                      | `I am ' happy '` → `I am 'happy'`            |
| **Article Fix**       | Change `a` → `an` before vowels/h                | `a honest man` → `an honest man`             |

---

## 3. Architecture Comparison

### 3.1 Pipeline

**Concept:** Sequential stages: tokenize → fix numbers → fix case → fix articles → punctuation → join.

* ✅ Easy to implement and extend.
* ❌ Struggles with deeply nested context.

### 3.2 Finite State Machine (FSM)

**Concept:** Parsing is modeled as states and transitions (e.g. `NORMAL → MARKER_FOUND → APPLY_RULE`).

* ✅ Strong contextual handling (quotes, punctuation clusters).
* ❌ More complex and harder to maintain.

### 3.3 Hybrid (Pipeline + FSM)

**Concept:** FSM handles tokenization & marker recognition; pipeline applies transformations.

* ✅ Combines modularity with context-awareness.
* ❌ Slightly more setup overhead.

**Decision:** All three will be implemented and benchmarked.

---

## 4. Rule Execution Order (Pipeline Flow)

1. Tokenizer / FSM Preprocessor
2. NumericFixer (hex, bin)
3. CaseFixer (up, low, cap)
4. ArticleFixer (a → an)
5. QuoteCleaner
6. PunctuationFixer
7. Joiner

---

## 5. Golden Test Set

| ID  | Input                                                                      | Expected Output                                                          |
| --- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
| T1  | `1E (hex) files were added`                                                | `30 files were added`                                                    |
| T2  | `It has been 10 (bin) years`                                               | `It has been 2 years`                                                    |
| T3  | `Ready, set, go (up) !`                                                    | `Ready, set, GO!`                                                        |
| T4  | `I should stop SHOUTING (low)`                                             | `I should stop shouting`                                                 |
| T5  | `Welcome to the brooklyn bridge (cap)`                                     | `Welcome to the Brooklyn Bridge`                                         |
| T6  | `This is so exciting (up, 2)`                                              | `This is SO EXCITING`                                                    |
| T7  | `I was sitting over there ,and then BAMM !!`                               | `I was sitting over there, and then BAMM!!`                              |
| T8  | `There it was. A amazing rock!`                                            | `There it was. An amazing rock!`                                         |
| T9  | `I am exactly how they describe me: ' awesome '`                           | `I am exactly how they describe me: 'awesome'`                           |
| T10 | `As Elton John said: ' I am the most well-known homosexual in the world '` | `As Elton John said: 'I am the most well-known homosexual in the world'` |

---

## 6. Tricky Test Cases

| ID | Input                      | Expected Output          | Description               |
| -- | -------------------------- | ------------------------ | ------------------------- |
| C1 | `a honest man`             | `an honest man`          | Edge case for “h”         |
| C2 | `10 (bin) and 1A (hex)`    | `2 and 26`               | Mixed numeric conversions |
| C3 | `HELLO (low, 2) WORLD`     | `hello world WORLD`      | Multi-word case           |
| C4 | `I waited ... and then ?!` | `I waited... and then?!` | Mixed punctuation         |
| C5 | `He said ' hello there '`  | `He said 'hello there'`  | Multi-word quotes         |

---

## 7. Large Paragraph Example

### Input

```
A friend sent me 1E (hex) messages yesterday , and I replied 10 (bin) times ! 
He said ' thanks ' but then shouted HELLO (low) . 
It was A honest mistake , I guess . 
This is truly amazing (up, 2) experience (cap) . 
Later that night , I wrote ' I am happy ' in my notebook (up) . 
Then I realized it was just a dream , a illusion that felt real . 
We talked about a orange , a apple , and a umbrella — all while laughing (cap, 4) . 
Finally , before I slept , I whispered ' good night ' and turned off the lights . 
It was a peaceful moment ... but also a reminder of how amazing (up, 3) everything (cap) can be .
```

### Expected Output

```
An friend sent me 30 messages yesterday, and I replied 2 times! 
He said 'thanks' but then shouted hello. 
It was an honest mistake, I guess. 
This is TRULY AMAZING Experience. 
Later that night, I wrote 'I am happy' in my NOTEBOOK. 
Then I realized it was just a dream, an illusion that felt real. 
We talked about an orange, an apple, and an umbrella — All While Laughing. 
Finally, before I slept, I whispered 'good night' and turned off the lights. 
It was a peaceful moment... but also a reminder of how AMAZING EVERYTHING Can Be.
```

---

## 8. CLI Mode Specification

The tool must be run from the terminal as follows:

```bash
go-reloaded <input_file> <output_file> <mode>
```

* **`<input_file>`** → the original text file containing markers.
* **`<output_file>`** → the corrected text after transformations.
* **`<mode>`** → processing mode, one of:

  * `pipeline` → sequential modular processor.
  * `fsm` → state machine processor.
  * `hybrid` → FSM-based tokenizer + pipeline transformations.

### Examples

```bash
# Run using pipeline architecture
go-reloaded input.txt output.txt pipeline

# Run using FSM architecture
go-reloaded input.txt output.txt fsm

# Run using hybrid architecture
go-reloaded input.txt output.txt hybrid
```

### Error Handling

* If arguments are missing → print usage help and exit with status `1`.
* If `<mode>` is invalid → print:

  ```
  Error: invalid mode. Use one of [pipeline|fsm|hybrid].
  ```
* If file I/O fails → exit with status `2`.

### Usage Help

```
Usage: go-reloaded <input_file> <output_file> <mode>
Modes:
  pipeline   Sequential modular processor
  fsm        Finite State Machine processor
  hybrid     FSM tokenizer + pipeline rules
```

---

## 9. Collaboration Model

* **Operator** → Issues tasks to the Agent via prompts.
* **Agent (GoReloadedAI)** → Implements tasks step by step using TDD.
* **Auditors** → Validate code correctness and test coverage.

Agents follow the flow described in `AGENTS.md`:
**Analyze → Write Tests → Implement Code → Refactor & Validate**

---

## 10. Conclusion

* **Pipeline** = baseline modular solution.
* **FSM** = better for context-sensitive parsing.
* **Hybrid** = production-ready balance.

All three implementations will be delivered and compared.

---

## 11. Project File Structure

Inspired by **Zone01’s governance style**, but simplified for the Go-Reloaded CLI project.
The structure balances **clarity, TDD workflows, and scalability**.

```
go-reloaded/
├── AGENTS.md            # Collaboration model (Operator, Agent, Auditor)
├── CONTRIBUTING.md      # How to contribute
├── RELEASE.md           # Release/benchmark process
├── LICENSE
├── README.md
│
├── cmd/
│   └── go-reloaded/
│       └── main.go      # CLI entry point
│
├── internal/
│   ├── processor/
│   │   ├── pipeline.go
│   │   ├── fsm.go
│   │   ├── hybrid.go
│   │   └── processor.go # Interface
│   │
│   ├── rules/
│   │   ├── numbers.go       # (hex), (bin)
│   │   ├── cases.go         # (up), (low), (cap)
│   │   ├── articles.go      # Article fixer
│   │   ├── quotes.go        # Quote cleaner
│   │   ├── punctuation.go   # Punctuation fixes
│   │   └── utils.go
│   │
│   └── tokenizer/
│       └── tokenizer.go
│
├── docs/
│   ├── Analysis.md      # This document (rules & specs)
│   ├── ARCHITECTURE.md  # Deep dive into pipeline, FSM, hybrid
│   ├── TESTING.md       # Golden/Tricky sets, TDD strategy
│   └── ROADMAP.md       # Future extensions and TODOs
│
├── tasks/
│   ├── TASK-001-cli-basics.md
│   ├── TASK-002-pipeline-core.md
│   ├── TASK-003-rule-hex-bin.md
│   ├── TASK-004-rule-case-formatting.md
│   ├── TASK-005-rule-articles.md
│   ├── TASK-006-fsm-core.md
│   └── TASK-007-hybrid-core.md
│
├── tests/
│   ├── golden_test.go
│   ├── tricky_test.go
│   ├── paragraph_test.go
│   └── cli_test.go
│
├── go.mod
└── go.sum
```

### 📌 Key Notes

* **Root = governance** → `AGENTS.md`, `CONTRIBUTING.md`, `RELEASE.md`.
* **`docs/` = specs** → design, architecture, testing, roadmap.
* **`tasks/` = TDD dev log** → one file per major milestone.
* **`internal/` = actual code** → rules, processors, tokenizer.
* **`tests/` = coverage** → Golden, Tricky, Paragraph, CLI.

---

## 12. Future Extensions

* `(rev)` → Reverse previous word(s).
* `(rep, n)` → Repeat previous word n times.
* Nested markers: `word (up) (hex)`.
* Configurable rule toggling via JSON/YAML.
* Multi-language grammar support.
* Performance benchmarking.
