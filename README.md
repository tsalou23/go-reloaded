# Go-Reloaded

A command-line text processing tool that applies various transformation rules to input text.

## Features

- **Number conversions**: Hexadecimal and binary to decimal
- **Case transformations**: Uppercase, lowercase, capitalization (single and multi-word)
- **Article corrections**: "a" → "an" before vowels and silent h
- **Quote cleaning**: Remove unnecessary spaces inside single quotes
- **Punctuation fixes**: Proper spacing around punctuation marks

## Usage

```bash
go-reloaded <input_file> <output_file> <mode>
```

### Modes

- `pipeline` - Sequential modular processor
- `fsm` - Finite State Machine processor  
- `hybrid` - FSM tokenizer + pipeline rules

### Examples

```bash
# Build the project
go build -o go-reloaded ./cmd/go-reloaded

# Run with pipeline mode
./go-reloaded input.txt output.txt pipeline

# Run with FSM mode
./go-reloaded input.txt output.txt fsm

# Run with hybrid mode
./go-reloaded input.txt output.txt hybrid
```

## Rule Examples

| Rule | Input | Output |
|------|-------|--------|
| (hex) | `1E (hex) files` | `30 files` |
| (bin) | `10 (bin) years` | `2 years` |
| (up) | `go (up)` | `GO` |
| (low) | `LOUD (low)` | `loud` |
| (cap) | `bridge (cap)` | `Bridge` |
| (up, 2) | `so exciting (up, 2)` | `SO EXCITING` |
| Articles | `a honest man` | `an honest man` |
| Quotes | `' hello '` | `'hello'` |
| Punctuation | `Hi , world !` | `Hi, world!` |

## Testing

```bash
# Run all tests
go test ./tests/ -v

# Run specific test suites
go test ./tests/golden_test.go -v
go test ./tests/tricky_test.go -v
go test ./tests/paragraph_test.go -v
go test ./tests/cli_test.go -v
```

## Project Structure

```
go-reloaded/
├── cmd/go-reloaded/     # CLI entry point
├── internal/
│   ├── processor/       # Pipeline, FSM, Hybrid processors
│   └── rules/          # Individual transformation rules
├── tests/              # Test suites
├── tasks/              # Development task tracking
└── docs/               # Documentation
```

## Development

This project follows Test-Driven Development (TDD) with tasks tracked in the `/tasks` directory. Each feature is implemented with comprehensive test coverage including Golden test cases and Tricky edge cases.

## Status

✅ All core tasks completed:
- TASK-001: CLI Basics
- TASK-002: Pipeline Core  
- TASK-003: Hex & Bin Rules
- TASK-004: Case Formatting Rules
- TASK-005: Article Rules
- TASK-006: Quote Rules
- TASK-007: Punctuation Rules

All Golden tests (T1-T10) and Tricky tests (C1-C5) are passing.