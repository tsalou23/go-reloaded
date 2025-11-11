# Go-Reloaded Auditor Testing Guide

*Simple testing guide for auditors to verify the text processing tool*

## 🚀 Getting Started

### 1. Setup
```bash
# Navigate to project directory
cd go-reloaded

# Build the main program
go build -o go-reloaded ./cmd/go-reloaded

# Build the demo (optional)
go build -o realtime-demo ./cmd/realtime-demo
```

### 2. Basic Usage
```bash
# Main program syntax:
./go-reloaded <input_file> <output_file> <mode>

# Available modes: pipeline, fsm, hybrid
```

## 📝 Quick Test (5 minutes)

### Step 1: Create a test file
```bash
echo "hello (up) world and FF (hex) items" > test.txt
```

### Step 2: Test all three modes
```bash
# Test Pipeline mode
./go-reloaded test.txt out1.txt pipeline
cat out1.txt
# Expected: HELLO world and 255 items

# Test FSM mode
./go-reloaded test.txt out2.txt fsm
cat out2.txt
# Expected: HELLO world and 255 items

# Test Hybrid mode
./go-reloaded test.txt out3.txt hybrid
cat out3.txt
# Expected: HELLO world and 255 items
```

### Step 3: Verify they're identical
```bash
diff out1.txt out2.txt && echo "Pipeline and FSM match!"
diff out1.txt out3.txt && echo "All modes match!"
```

### Step 4: Test real-time demo (optional)
```bash
./realtime-demo
# Type: hello(up)
# Expected output: HELLO
# Type: quit to exit
```

## 🔍 Comprehensive Testing

### Test All Transformation Rules

```bash
# Create comprehensive test file
cat > full_test.txt << 'EOF'
1E (hex) files and 10 (bin) users
hello (up) world and LOUD (low) text
bridge (cap) over water
so exciting (up, 2) news today
a apple and a honest man
He said ' hello , world ! '
EOF

# Test with pipeline mode
./go-reloaded full_test.txt result.txt pipeline
cat result.txt
```

**Expected Output:**
```
30 files and 2 users
HELLO world and loud text
Bridge over water
SO EXCITING news today
an apple and an honest man
He said 'hello, world!'
```

### Common Usage Patterns

```bash
# Process a file with pipeline mode (recommended)
./go-reloaded input.txt output.txt pipeline

# Use FSM mode for large files (memory efficient)
./go-reloaded large_file.txt output.txt fsm

# Use hybrid mode for complex preprocessing
./go-reloaded input.txt output.txt hybrid

# Interactive testing with real-time demo
./realtime-demo
```

## 📋 What Each Rule Does

| Rule | Input Example | Output Example |
|------|---------------|----------------|
| Hex | `FF (hex) items` | `255 items` |
| Binary | `1010 (bin) users` | `10 users` |
| Uppercase | `hello (up) world` | `HELLO world` |
| Lowercase | `LOUD (low) noise` | `loud noise` |
| Capitalize | `bridge (cap) water` | `Bridge water` |
| Multi-word | `exciting (up, 2) news` | `EXCITING NEWS` |
| Articles | `a apple` | `an apple` |
| Quotes | `' hello , world ! '` | `'hello, world!'` |

## 🚀 Simple Test Commands

```bash
# Test hex conversion
echo "FF (hex) items" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt  # Should show: 255 items

# Test case transformation
echo "hello (up) world" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt  # Should show: HELLO world

# Test article correction
echo "a apple" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt  # Should show: an apple

# Test quote cleaning
echo "He said ' hello , world ! '" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt  # Should show: He said 'hello, world!'
```

## 🏗️ Architecture Verification

### Memory Usage Test (Large Files)

```bash
# Create large test file
for i in {1..1000}; do echo "hello (up) world and FF (hex) items"; done > large_test.txt

# Test memory usage with different modes
time ./go-reloaded large_test.txt large_out_pipeline.txt pipeline
time ./go-reloaded large_test.txt large_out_fsm.txt fsm  
time ./go-reloaded large_test.txt large_out_hybrid.txt hybrid

# FSM should use least memory, Pipeline should be fastest
```

### Error Handling Test

```bash
# Test invalid arguments
./go-reloaded  # Should show usage
./go-reloaded input.txt  # Should show usage
./go-reloaded input.txt output.txt invalid_mode  # Should show error

# Test missing input file
./go-reloaded nonexistent.txt output.txt pipeline  # Should show error
```

## 🧪 Automated Test Suite

```bash
# Run comprehensive test suite
go test ./tests/ -v

# Run specific test categories
go test ./tests/golden_test.go -v
go test ./tests/tricky_test.go -v
go test ./tests/cli_test.go -v
```

**Expected Result:** All tests should PASS

## 🔧 Edge Cases Testing

### Complex Combinations

```bash
cat > complex_test.txt << 'EOF'
a FF (hex) items and 1010 (bin) users said ' hello (up) world ! '
The VERY LOUD (low, 2) noise from a honest (cap) man .
EOF

./go-reloaded complex_test.txt complex_out.txt pipeline
cat complex_out.txt
```

**Expected Output:**
```
an 255 items and 10 users said 'HELLO world!'
The very loud noise from an Honest man.
```

## 📊 Performance Benchmarking

### Speed Comparison

```bash
# Create benchmark file
for i in {1..100}; do echo "hello (up) world FF (hex) items 1010 (bin) users"; done > benchmark.txt

# Time each mode
time ./go-reloaded benchmark.txt bench_pipeline.txt pipeline
time ./go-reloaded benchmark.txt bench_fsm.txt fsm
time ./go-reloaded benchmark.txt bench_hybrid.txt hybrid
```

**Expected Performance Order:**
1. Pipeline (fastest)
2. FSM (moderate)
3. Hybrid (slowest due to tokenization overhead)

## ✅ Validation Checklist

### Functional Requirements
- [ ] All transformation rules work correctly
- [ ] All three modes produce identical results
- [ ] Error handling works properly
- [ ] File I/O operations work correctly

### Non-Functional Requirements  
- [ ] FSM mode uses minimal memory
- [ ] Pipeline mode has good performance
- [ ] Real-time mode responds immediately
- [ ] All modes handle large files

### Code Quality
- [ ] All tests pass (42 test cases)
- [ ] No runtime errors or panics
- [ ] Consistent output across modes
- [ ] Proper error messages

## 🚨 Common Issues to Check

### 1. Rule Processing Order
Articles should be applied LAST to avoid conflicts:
```bash
# This should work correctly
echo "a FF (hex) items" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt
# Expected: "an 255 items" (not "a 255 items")
```

### 2. Case Preservation
Already uppercase words shouldn't be changed by (cap):
```bash
echo "NASA (cap) agency" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt
# Expected: "NASA agency" (not "Nasa agency")
```

### 3. Multi-word Transformations
Should affect the correct number of preceding words:
```bash
echo "hello beautiful world (up, 2)" > test.txt
./go-reloaded test.txt out.txt pipeline
cat out.txt
# Expected: "hello BEAUTIFUL WORLD"
```

## 📝 Test Report Template

```
Go-Reloaded Audit Report
========================

Date: ___________
Auditor: ___________

Basic Functionality:
[ ] Pipeline mode: PASS/FAIL
[ ] FSM mode: PASS/FAIL  
[ ] Hybrid mode: PASS/FAIL
[ ] Real-time demo: PASS/FAIL

Transformation Rules:
[ ] Number conversions: PASS/FAIL
[ ] Case transformations: PASS/FAIL
[ ] Article corrections: PASS/FAIL
[ ] Quote/punctuation: PASS/FAIL

Architecture Verification:
[ ] Identical outputs: PASS/FAIL
[ ] Memory efficiency: PASS/FAIL
[ ] Error handling: PASS/FAIL
[ ] Test suite: PASS/FAIL

Overall Assessment: PASS/FAIL
Notes: ___________
```

## 🎯 Quick Verification Commands

```bash
# One-liner to test all modes with same input
echo "hello (up) FF (hex) items" > quick_test.txt
./go-reloaded quick_test.txt out1.txt pipeline && \
./go-reloaded quick_test.txt out2.txt fsm && \
./go-reloaded quick_test.txt out3.txt hybrid && \
echo "Pipeline:" && cat out1.txt && \
echo "FSM:" && cat out2.txt && \
echo "Hybrid:" && cat out3.txt && \
echo "All identical:" && diff out1.txt out2.txt && diff out1.txt out3.txt && echo "YES"
```

**Expected:** All outputs should be identical: `HELLO 255 items`

## 🔍 Real-time Demo Notes

The real-time demo (`./realtime-demo`) supports basic transformations only:
- ✅ Single word transformations: `hello(up)` → `HELLO`
- ✅ Number conversions: `FF(hex)` → `255`
- ✅ Article corrections: `a apple` → `an apple`
- ❌ Multi-word transformations: `(up, 2)` not supported

For full functionality testing, use the main CLI with pipeline/fsm/hybrid modes.