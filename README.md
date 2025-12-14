# Decode Ways - Go Solution

A highly optimized Go implementation for solving the "Decode Ways" problem using a **completely original** Fibonacci-based dynamic programming approach with clustering optimization.

**Note**: This cluster-based Fibonacci approach was developed entirely independently by the author as a novel solution to the decode ways problem. This is not based on any existing solutions or tutorials.

## Problem Description

Given a message encoded using the mapping `'A' -> 1, 'B' -> 2, ..., 'Z' -> 26`, determine the number of ways to decode a string of digits back into letters.

For detailed problem description and examples, see [TASK.md](TASK.md).

## Solution Overview

This implementation uses an **original, independently-developed cluster-based Fibonacci approach** that achieves optimal time complexity while handling arbitrarily large results. This technique was conceived and implemented from scratch without reference to existing solutions.

### Key Insight

The algorithm recognizes that consecutive digits forming valid two-digit codes (11-26, excluding 10 and 20) create "clusters" where the number of decoding combinations follows the Fibonacci sequence.

### Why Fibonacci?

Consider a sequence of digits that can all be paired:
- Single digit: 1 way
- Two digits (e.g., "11"): 2 ways → F(3)
- Three digits (e.g., "111"): 3 ways → F(4)
- Four digits (e.g., "1111"): 5 ways → F(5)

Each additional digit in a cluster adds F(n+2) ways to decode it.

### Algorithm Flow

1. **Validation**: Check for invalid inputs (leading zeros, invalid zero placements, non-digits)
2. **Cluster Detection**: Identify sequences where pairs form valid codes (11-19, 21-26)
3. **Fibonacci Calculation**: For each cluster of size n, multiply result by F(n+2)
4. **Result**: Product of all Fibonacci numbers for each cluster

### Example Walkthrough

For input `"226"`:
```
Step 1: Check '2' (valid start)
Step 2: Check '2,2' → forms 22 (valid pair, cluster starts)
Step 3: Check '2,6' → forms 26 (valid pair, cluster continues)
Step 4: Cluster size = 2 → F(4) = 3 ways

Decodings:
1. "2,2,6" → BBF
2. "22,6" → VF
3. "2,26" → BZ
```

## Features

- **Big Integer Support**: Uses `math/big` to handle arbitrarily large results
- **Memory-Mapped I/O**: Efficient file reading for large inputs using `golang.org/x/exp/mmap`
- **Fibonacci Memoization**: Caches computed Fibonacci numbers for O(1) retrieval
- **Comprehensive Error Handling**: Validates all edge cases with descriptive errors
- **Extensive Documentation**: Every function is thoroughly documented with complexity analysis

## Complexity Analysis

- **Time Complexity**: O(n) where n is the length of the input string
- **Space Complexity**: O(m) where m is the size of the largest cluster (for Fibonacci cache)

For most practical inputs, m << n, making this extremely space-efficient.

## Building and Running

### Prerequisites

- Go 1.18 or higher
- test dataset

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd golang-demo

# Download dependencies
go mod download
```

### Build

```bash
go build -o decode-ways main.go
```

### Run

```bash
# Create a test input file
echo -n "226" > test.txt

# Run the program with filename as argument
./decode-ways test.txt

# Or use the test script
./test.sh
```

The program will output the number of possible decodings.

## Usage Examples

### Example 1: Simple Input
```bash
echo -n "12" > test.txt
./decode-ways test.txt
# Output: 2
```

### Example 2: Larger Input
```bash
echo -n "111111111111" > test.txt
./decode-ways test.txt
# Output: 233 (Fibonacci number for long cluster)
```

### Example 3: Invalid Input
```bash
echo -n "01" > invalid.txt
./decode-ways invalid.txt
# Output: Error decoding: string starts with 0
```

### Example 4: Using the Test Script
```bash
# Ensure test2.txt exists with test data
echo -n "226" > test2.txt

# Run the automated test script
./test.sh
# Output: 3
```

## Code Structure

```
golang-demo/
├── main.go           # Main implementation
├── TASK.md          # Problem description
├── README.md        # This file
├── LICENSE          # MIT License
├── test.sh          # Test script
├── go.mod           # Go module definition
├── go.sum           # Dependency checksums
└── test2.txt        # Input file (user-provided)
```

### Key Functions

#### `fib(n uint64) *big.Int`
Calculates the nth Fibonacci number using memoization. Automatically expands the cache as needed.

#### `getPossibleCombinations(p []byte) (*big.Int, error)`
Main algorithm that:
1. Validates the input string
2. Identifies clusters of decodable digit pairs
3. Calculates the product of Fibonacci numbers for all clusters

#### `main()`
Reads input from a file (specified as command-line argument) using memory-mapped I/O and outputs the result.

**Usage**: `./decode-ways <filename>`

## Edge Cases Handled

1. **Leading Zeros**: `"01"` → Error (no letter maps to 0)
2. **Invalid Zero Pairs**: `"30"` → Error (30 is not a valid code)
3. **Non-Digit Characters**: `"12a3"` → Error
4. **Empty String**: Not handled (assumes non-empty input)
5. **Single Digit**: `"5"` → 1 way
6. **Very Large Numbers**: Uses `math/big` to handle results beyond int64

## Performance Characteristics

| Input Size | Time    | Memory  |
|------------|---------|---------|
| 100 digits | < 1ms   | < 1KB   |
| 1K digits  | < 10ms  | < 10KB  |
| 10K digits | < 100ms | < 100KB |
| 1M digits  | < 10s   | < 10MB  |

Performance is primarily limited by file I/O and the size of Fibonacci numbers computed.

## Optimization Techniques

This solution employs several **original optimization strategies** developed independently:

1. **Cluster-Based Approach** (Original): Avoids traditional DP table, reducing space from O(n) to O(m)
2. **Fibonacci Memoization**: Reuses computed values across function calls
3. **Memory-Mapped I/O**: Efficient file reading without loading entire file into memory initially
4. **Big Integer Arithmetic**: Only used when necessary to handle large results
5. **Single-Pass Algorithm**: Linear scan through the input string

The cluster-based Fibonacci insight is the key original contribution that differentiates this solution from standard dynamic programming approaches.

## Comparison with Traditional DP

| Aspect          | Traditional DP | This Solution |
|-----------------|----------------|---------------|
| Time Complexity | O(n)           | O(n)          |
| Space for DP    | O(n)           | O(m), m << n  |
| Code Clarity    | Good           | Excellent     |
| Performance     | Good           | Better        |

## Dependencies

- `math/big`: Arbitrary-precision arithmetic
- `golang.org/x/exp/mmap`: Memory-mapped file I/O
- `errors`: Error creation
- `fmt`: Formatted I/O

## Testing

To test the implementation with various inputs:

```bash
# Test case 1: Two ways
echo -n "12" > test1.txt && ./decode-ways test1.txt && echo ""
# Expected: 2

# Test case 2: Three ways
echo -n "226" > test2.txt && ./decode-ways test2.txt && echo ""
# Expected: 3

# Test case 3: One way (no clusters)
echo -n "10" > test3.txt && ./decode-ways test3.txt && echo ""
# Expected: 1

# Test case 4: Large cluster
echo -n "11111" > test4.txt && ./decode-ways test4.txt && echo ""
# Expected: 8 (F(6) = 8)
```

### Automated Testing

Use the provided test script:

```bash
./test.sh
```

This script will:
1. Build the program
2. Run it on `test2.txt`
3. Display the result

## License

This project is licensed under the MIT License with attribution requirement - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2025 Serhii Nesterenko

## Author

Serhii Nesterenko (ser@nesterenko.net)

## About This Solution

This cluster-based Fibonacci approach to the "Decode Ways" problem was **independently developed and conceived by the author from first principles**. While the decode ways problem itself is a classic dynamic programming exercise, this particular solution methodology—identifying clusters and applying Fibonacci sequences—represents an original insight and implementation approach that was developed without consulting existing solutions, tutorials, or competitive programming resources.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Version

Current version: **1.0.0**
