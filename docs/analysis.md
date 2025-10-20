# 🧠 TextFixer — Analysis Document

## 📑 Table of Contents

1. Problem Description
2. Rule Catalog
3. Architecture Comparison
    3.1 Pipeline
    3.2 Finite State Machine (FSM)
    3.3 Hybrid (Pipeline + FSM)
4. Rule Execution Order
5. Golden Test Set
6. Tricky Test Cases
7. Large Paragraph Example
8. Conclusion
9. Future Extensions

---

## 1. Problem Description

**TextFixer** (also called *Go-reloaded*) is a command-line tool written in Go that reads an input text file, detects special markers such as `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)`, etc., applies the corresponding transformations, and produces a **cleaned and corrected output file**.

The tool must correctly handle:

* **Number conversions:** hexadecimal and binary → decimal
* **Word formatting:** uppercase, lowercase, capitalization
* **Punctuation corrections:** proper spacing and structure
* **Grammar fixes:** article corrections (“a” → “an”)
* **Quotes handling:** proper spacing inside `' '`
* **Combined or nested patterns:** multi-word markers like `(up, 3)`

The analysis explores **three architectures** to process text transformations:

1. **Pipeline:** Sequential stages; each performs a transformation.
2. **Finite State Machine (FSM):** Rule-based state transitions.
3. **Hybrid:** Combination of FSM (for contextual parsing) + Pipeline (for transformations).

---

## 2. Rule Catalog

| Rule                  | Description                                       | Example                                      |
| --------------------- | ------------------------------------------------- | -------------------------------------------- |
| `(hex)`               | Convert previous word from hexadecimal to decimal | `1E (hex)` → `30`                            |
| `(bin)`               | Convert previous word from binary to decimal      | `10 (bin)` → `2`                             |
| `(up)`                | Convert previous word to uppercase                | `go (up)` → `GO`                             |
| `(low)`               | Convert previous word to lowercase                | `LOUD (low)` → `loud`                        |
| `(cap)`               | Capitalize previous word                          | `bridge (cap)` → `Bridge`                    |
| `(up, n)`             | Convert previous *n* words to uppercase           | `so exciting (up, 2)` → `SO EXCITING`        |
| `(low, n)`            | Convert previous *n* words to lowercase           | `HELLO WORLD (low, 2)` → `hello world`       |
| `(cap, n)`            | Capitalize previous *n* words                     | `the big bridge (cap, 3)` → `The Big Bridge` |
| **Punctuation**       | Fix spaces before punctuation                     | `Hi , world !` → `Hi, world!`                |
| **Multi-punctuation** | Keep `!?` or `...` sequences intact               | `Wait ... What ?` → `Wait... What?`          |
| **Quotes `' '`**      | Remove spaces inside quotes                       | `I am ' happy '` → `I am 'happy'`            |
| **Article Fix**       | Replace “a” with “an” before vowels/h             | `a honest man` → `an honest man`             |

---

## 3. Architecture Comparison

### 3.1 Pipeline Architecture

**Concept:**
Each transformation rule runs in sequence through a “pipeline.”
Each stage receives text, modifies it, and passes it to the next.

**Flow Example:**

```
Input → Tokenizer → NumericFixer → CaseFixer → ArticleFixer → QuoteCleaner → PunctuationFixer → Output
```

**Advantages:**

* Simple, modular, and easy to extend.
* Each rule can be tested individually.
* Debug-friendly (inspect intermediate outputs).

**Disadvantages:**

* Difficult to handle complex or nested markers.
* Sequential only — no backtracking or parallelism.

---

### 3.2 Finite State Machine (FSM) Architecture

**Concept:**
The text processor acts as a **state machine**.
Each state represents a parsing context (word, punctuation, marker, inside quotes, etc.).
Transitions occur when the parser encounters markers or punctuation.

**Example States:**

```
NORMAL → MARKER_FOUND → APPLY_RULE → RETURN_NORMAL
QUOTE_OPEN → INSIDE_QUOTE → QUOTE_CLOSE
```

**Advantages:**

* Excellent for context-aware processing.
* Handles nested or dependent rules effectively.
* Suitable for grammatically complex text.

**Disadvantages:**

* More complex to design and debug.
* Harder to maintain when adding new rules.
* Requires a solid state transition map.

**Best For:**
Tools needing precise parsing and nested structure awareness.

---

### 3.3 Hybrid Architecture (Pipeline + FSM)

**Concept:**
Combines both worlds:

* **FSM:** Handles tokenization, markers, and contextual parsing (quotes, nested structures).
* **Pipeline:** Executes transformations in clean, isolated stages.

**Flow Example:**

```
Input → FSM Tokenizer (handles quotes & markers) → Pipeline (rules in sequence) → Output
```

**Advantages:**

* Balances modularity and context-awareness.
* Easier to debug than a pure FSM.
* Flexible enough for production-scale systems.

**Disadvantages:**

* Slightly more setup and coordination between FSM & Pipeline.

**Best For:**
Production-level systems where both context and modularity are required.

---

## 4. Rule Execution Order (Pipeline Flow)

| Step | Stage                            | Description                               |
| ---- | -------------------------------- | ----------------------------------------- |
| 1    | **Tokenizer / FSM Preprocessor** | Splits text, detects quotes and markers   |
| 2    | **NumericFixer**                 | Converts `(hex)` and `(bin)`              |
| 3    | **CaseFixer**                    | Applies `(up)`, `(low)`, `(cap)`          |
| 4    | **ArticleFixer**                 | Fixes `a` → `an`                          |
| 5    | **QuoteCleaner**                 | Removes extra spaces inside `' '`         |
| 6    | **PunctuationFixer**             | Adjusts punctuation spacing               |
| 7    | **Joiner**                       | Combines final cleaned tokens into output |

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

| ID | Input                      | Expected Output          | Description                   |
| -- | -------------------------- | ------------------------ | ----------------------------- |
| C1 | `a honest man`             | `an honest man`          | Edge case for h-starting word |
| C2 | `10 (bin) and 1A (hex)`    | `2 and 26`               | Mixed numeric conversions     |
| C3 | `HELLO (low, 2) WORLD`     | `hello world WORLD`      | Multi-word case               |
| C4 | `I waited ... and then ?!` | `I waited... and then?!` | Mixed punctuation             |
| C5 | `He said ' hello there '`  | `He said 'hello there'`  | Multi-word quotes             |

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

## 8. Conclusion

| Architecture | Pros                           | Cons                           | Best Use                                            |
| ------------ | ------------------------------ | ------------------------------ | --------------------------------------------------- |
| **Pipeline** | Simple, modular, testable      | Not ideal for nested context   | Sequential rule transformations                     |
| **FSM**      | Context-aware, precise parsing | Complex to maintain and extend | Deep contextual text parsing                        |
| **Hybrid**   | Balanced, practical, flexible  | Slightly higher complexity     | Real-world systems combining context and modularity |

**Final Decision:**
All three architectures will be implemented for experimentation and benchmarking:

* **Pipeline:** Baseline modular version
* **FSM:** Research-oriented contextual processor
* **Hybrid:** Production-ready architecture (FSM + Pipeline)

---

## 9. Future Extensions

* Add `(rev)` → Reverse previous word(s)
* Add `(rep, n)` → Repeat previous word *n* times
* Support nested markers (e.g., `word (up) (hex)`)
* Configurable YAML/JSON rule toggling
* Multi-language grammar adjustments
* Benchmark performance between Pipeline, FSM, and Hybrid implementations

---

**Author:** Giorgos Tsaloukidis
**Date:** October 2025
**Language:** English
**Architectures:** Pipeline • FSM • Hybrid
**Project:** TextFixer (Go-Reloaded)
**Goal:** Build three working text processors and compare modularity, readability, and runtime efficiency.
