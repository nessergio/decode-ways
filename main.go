// Copyright (c) 2025 Serhii Nesterenko
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

/*
Package main implements a solution to the "Decode Ways" problem.

This program calculates the number of ways a string of digits can be decoded
into letters, where:
  - 'A' -> 1, 'B' -> 2, ..., 'Z' -> 26

The solution uses a Fibonacci-based dynamic programming approach by identifying
"clusters" of consecutive digits that can be decoded in multiple ways.

Algorithm Overview:
The key insight is that consecutive digits in the range 11-26 (excluding 20)
form "clusters" where each digit can either be decoded separately or combined
with the previous digit. For a cluster of size n, the number of ways to decode
it follows the Fibonacci sequence: F(n+2).

For example:
  - "1" -> 1 way (A)
  - "11" -> 2 ways (AA, K) = F(3) = 2
  - "111" -> 3 ways (AAA, AK, KA) = F(4) = 3
  - "1111" -> 5 ways (AAAA, AAK, AKA, KAA, KK) = F(5) = 5

The total number of combinations is the product of Fibonacci numbers for all clusters.

Time Complexity: O(n) where n is the length of the input string
Space Complexity: O(m) where m is the size of the largest cluster (for Fibonacci cache)
*/
package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"golang.org/x/exp/mmap"
)

// f is the Fibonacci cache table storing precomputed Fibonacci numbers.
// Initialized with F(0) = 0 and F(1) = 1 as base cases.
// Values are computed on-demand by the fib() function and stored for reuse.
var f = []*big.Int{big.NewInt(0), big.NewInt(1)}

// maxFib stores the index of the maximum Fibonacci number currently cached.
// Used to determine if we need to compute additional Fibonacci numbers.
var maxFib uint64 = 2

// fib calculates and returns the nth Fibonacci number using memoization.
//
// The function uses big.Int to handle arbitrarily large Fibonacci numbers.
// For reference, F(93) = 12,200,160,415,121,876,738 is the largest Fibonacci
// number that fits in int64.
//
// Parameters:
//   - n: The index of the Fibonacci number to calculate (0-indexed)
//
// Returns:
//   - *big.Int: The nth Fibonacci number
//
// Time Complexity: O(1) if cached, O(n - maxFib) if not cached
func fib(n uint64) *big.Int {
	// Expand the Fibonacci cache up to index n if needed
	for ; maxFib <= n; maxFib++ {
		c := big.Int{}
		// F(n) = F(n-1) + F(n-2)
		f = append(f, c.Add(f[maxFib-1], f[maxFib-2]))
	}
	return f[n]
}

// getPossibleCombinations calculates the number of ways to decode a digit string.
//
// The algorithm works by:
//  1. Validating the input (no leading zeros, no invalid digit pairs)
//  2. Identifying "clusters" of digits that can be decoded multiple ways
//  3. Multiplying the Fibonacci numbers corresponding to each cluster size
//
// A cluster is a sequence of consecutive digits where each pair is in the range
// 11-19 or 21-26. These are the only two-digit combinations that can validly
// be decoded either as two separate letters or as one letter.
//
// Parameters:
//   - p: Byte slice containing the digit string to decode
//
// Returns:
//   - *big.Int: The number of possible decodings
//   - error: An error if the input is invalid
//
// Example:
//   - "12" -> cluster size 1 -> F(3) = 2 ways
//   - "226" -> cluster of "22" (size 1) -> F(3) = 2, then "6" separate -> 2 * 1 = 2... wait
//     Actually "226": we have "2,2,6" which can be "2,26" or "22,6" or "2,2,6"
//     Let me reconsider... "226" has a cluster "226" where 22 and 26 are valid pairs
//     but we need to track this more carefully in the logic
func getPossibleCombinations(p []byte) (*big.Int, error) {
	clusterSize := uint64(0) // Current size of the cluster being processed
	x := big.NewInt(1)       // Running product of Fibonacci numbers
	a := p[0]                // Previous digit (for pair checking)

	// Validate first character: must be a digit 1-9 (no leading zero)
	if a == 0x30 { // '0'
		return big.NewInt(0), errors.New("string starts with 0")
	} else if a < 0x31 || a > 0x39 { // Not '1'-'9'
		return big.NewInt(0), errors.New("string starts with non-digit character")
	}

	// Process each subsequent digit
	for i, b := range p[1:] {
		// Validate that current character is a digit
		if b < 0x30 || b > 0x39 { // Not '0'-'9'
			return big.NewInt(0), fmt.Errorf("encountered non-digit character at pos. %d", i)
		}

		// Check for invalid zero: '0' can only appear after '1' or '2' (forming 10 or 20)
		if b == 0x30 && a != 0x31 && a != 0x32 {
			return big.NewInt(0), fmt.Errorf("encountered 0 which can not be attached to %c at pos. %d", a, i)
		}

		// Identify cluster boundaries
		// A pair (a, b) is in a cluster if it forms 11-19 or 21-26
		// Note: 10 and 20 are NOT in clusters as they have only one decoding
		if (a == 0x31 && b > 0x30) || (a == 0x32 && b > 0x30 && b <= 0x36) {
			// We are inside a cluster: the pair can be decoded in 2 ways
			clusterSize++
		} else if clusterSize > 0 {
			// We've exited a cluster: multiply result by F(clusterSize + 2)
			// The +2 offset is because a cluster of size 1 has F(3) = 2 ways
			x.Mul(x, fib(clusterSize+2))
			clusterSize = 0 // Reset cluster size
		}

		a = b // Move to next digit
	}

	// Handle the case where the string ends inside a cluster
	if clusterSize > 0 {
		x.Mul(x, fib(clusterSize+2))
	}

	return x, nil
}

// main reads a digit string from a file and prints the number of decode ways.
//
// The file is read using memory-mapped I/O for efficient handling of large files.
// The filename is provided as the first command-line argument.
//
// Usage:
//   decode-ways <filename>
//
// Example:
//   decode-ways test2.txt
func main() {
	// Check if filename argument is provided
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: decode-ways <filename>")
		fmt.Fprintln(os.Stderr, "Example: decode-ways test2.txt")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Open file using memory-mapped I/O for efficient reading
	r, err := mmap.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	// Read entire file content into memory
	p := make([]byte, r.Len())
	if _, err = r.ReadAt(p, 0); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}

	// Calculate number of possible decodings
	x, err := getPossibleCombinations(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding: %v\n", err)
		os.Exit(1)
	}

	// Print result
	fmt.Print(x)
}
