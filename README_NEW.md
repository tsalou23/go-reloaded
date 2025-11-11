# Go-Reloaded

A high-performance command-line text processing tool that applies intelligent transformation rules to input text using multiple processing architectures.

## 🚀 Features

### Core Transformations
- **Number System Conversions**: Hexadecimal and binary to decimal conversion
- **Case Transformations**: Uppercase, lowercase, and capitalization (single and multi-word)
- **Article Corrections**: Automatic "a" → "an" corrections before vowels and silent h
- **Quote Normalization**: Intelligent spacing removal inside single quotes
- **Punctuation Optimization**: Proper spacing around punctuation marks

### Processing Architectures
- **Pipeline Processor**: Sequential rule-based processing with regex optimization
- **FSM Processor**: Pure finite state machine with character-by-character processing
- **Hybrid Processor**: FSM tokenization combined with pipeline rule application
- **Real-time FSM**: Live text transformation as you type

## 📦 Installation

```bash
# Clone the repository
git clone <repository-url>
cd go-reloaded

# Build the main application
go build -o go-reloaded ./cmd/go-reloaded

# Build the real-time demo
go build -o realtime-demo ./cmd/realtime-demo
```

## 🎯 Usage

### Command Line Interface

```bash
go-reloaded <input_file> <output_file> <mode>
```

**Available Modes:**
- `pipeline` - Sequential modular processor (fastest for small files)
- `fsm` - Finite State Machine processor (most memory efficient)
- `hybrid` - FSM tokenizer + pipeline rules (balanced approach)

### Real-time Processing

```bash
./realtime-demo
```

Interactive terminal interface for testing transformations in real-time.

## 📋 Transformation Rules

| Rule Type | Syntax | Input Example | Output Example |
|-----------|--------|---------------|----------------|
| Hexadecimal | `(hex)` | `1E (hex) files` | `30 files` |
| Binary | `(bin)` | `10 (bin) years` | `2 years` |
| Uppercase | `(up)` | `go (up)` | `GO` |
| Lowercase | `(low)` | `LOUD (low)` | `loud` |
| Capitalize | `(cap)` | `bridge (cap)` | `Bridge` |
| Multi-word Up | `(up, 2)` | `so exciting (up, 2)` | `SO EXCITING` |
| Multi-word Low | `(low, 3)` | `VERY LOUD NOISE (low, 3)` | `very loud noise` |
| Multi-word Cap | `(cap, 2)` | `new york (cap, 2)` | `New York` |
| Article Fix | Auto | `a honest man` | `an honest man` |
| Quote Clean | Auto | `' hello world '` | `'hello world'` |
| Punctuation | Auto | `Hi , world !` | `Hi, world!` |

## 🏗️ Architecture Comparison

| Feature | Pipeline | FSM | Hybrid |
|---------|----------|-----|--------|
| **Memory Usage** | Moderate | Minimal | High |
| **Processing Speed** | Fast | Moderate | Moderate |
| **Large File Handling** | Good | Excellent | Poor |
| **Code Complexity** | Low | Medium | High |
| **Best Use Case** | General purpose | Memory-constrained | Complex preprocessing |

## 🧪 Testing

### Comprehensive Test Suite

```bash
# Run all tests
go test ./tests/ -v

# Run specific test categories
go test ./tests/golden_test.go -v          # Standard test cases
go test ./tests/tricky_test.go -v          # Edge cases and complex scenarios
go test ./tests/paragraph_test.go -v       # Multi-paragraph processing
go test ./tests/cli_test.go -v             # Command-line interface tests
```

### Test Coverage
- **42 comprehensive test cases** covering all transformation rules
- **Golden test suite** with standard input/output pairs
- **Tricky test suite** for edge cases and rule interactions
- **100% pass rate** across all processors

## 📁 Project Structure

```
go-reloaded/
├── cmd/
│   ├── go-reloaded/         # Main CLI application
│   └── realtime-demo/       # Interactive real-time demo
├── internal/
│   ├── processor/           # Core processing engines
│   │   ├── pipeline.go      # Sequential rule processor
│   │   ├── fsm.go          # Finite state machine
│   │   ├── hybrid.go       # Hybrid tokenizer + rules
│   │   └── realtime_fsm.go # Real-time character processor
│   ├── rules/              # Transformation rule modules
│   │   ├── numbers.go      # Hex/binary conversions
│   │   ├── cases.go        # Case transformations
│   │   ├── articles.go     # Article corrections
│   │   ├── quotes.go       # Quote normalization
│   │   └── punctuation.go # Punctuation fixes
│   └── tokenizer/          # FSM-based tokenization
├── tests/                  # Comprehensive test suites
├── tasks/                  # Development task tracking
└── docs/                   # Documentation
```

## 🔧 Development

### Prerequisites
- Go 1.19 or higher
- Git for version control

### Development Workflow
This project follows **Test-Driven Development (TDD)** methodology:

1. **Task Definition**: Requirements tracked in `/tasks` directory
2. **Test Creation**: Comprehensive test cases written first
3. **Implementation**: Code written to pass all tests
4. **Validation**: All processors tested against identical test suites

### Adding New Rules
1. Create rule function in appropriate `/internal/rules/` file
2. Add comprehensive test cases in `/tests/`
3. Update all three processors to handle the new rule
4. Verify consistent behavior across all architectures

## 🎮 Interactive Examples

### Basic Transformations
```bash
# Input: "hello (up) world"
# Output: "HELLO world"

# Input: "FF (hex) items"
# Output: "255 items"

# Input: "1010 (bin) users"
# Output: "10 users"
```

### Complex Scenarios
```bash
# Input: "a apple and a orange"
# Output: "an apple and an orange"

# Input: "He said ' hello , world ! '"
# Output: "He said 'hello, world!'"

# Input: "process (cap, 2) this (up) text"
# Output: "PROCESS THIS THIS text"
```

### Real-time Demo
```bash
./realtime-demo
Input: hello(up)
Output: HELLO

Input: 1E(hex) files
Output: 30 files
```

## 📊 Performance Characteristics

### Memory Efficiency (Large Files)
1. **FSM**: O(1) memory usage - constant regardless of file size
2. **Pipeline**: O(n) memory usage - linear with file size
3. **Hybrid**: O(2n) memory usage - stores tokens + original text

### Processing Speed (Small Files)
1. **Pipeline**: Fastest - optimized regex processing
2. **Hybrid**: Moderate - tokenization overhead
3. **FSM**: Slower - character-by-character processing

## ✅ Status

**All core development tasks completed:**
- ✅ TASK-001: CLI Foundation & File I/O
- ✅ TASK-002: Pipeline Architecture & Core Rules
- ✅ TASK-003: Number System Conversions (Hex/Binary)
- ✅ TASK-004: Case Transformation Rules
- ✅ TASK-005: Article Correction System
- ✅ TASK-006: Quote Normalization
- ✅ TASK-007: Punctuation Optimization
- ✅ **Bonus**: Real-time FSM Implementation

**Test Results:**
- 🎯 All Golden tests (T1-T10): **PASSING**
- 🎯 All Tricky tests (C1-C5): **PASSING**
- 🎯 42 Comprehensive test cases: **100% PASS RATE**

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write comprehensive tests for your changes
4. Implement your feature following TDD principles
5. Ensure all tests pass across all processors
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ using Test-Driven Development and Go**