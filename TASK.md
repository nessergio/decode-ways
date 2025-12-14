# Decode Ways Problem

## Problem Statement

You are given a message encoded using the following mapping:

```
'A' -> 1
'B' -> 2
'C' -> 3
...
'Z' -> 26
```

Write a function or algorithm that takes a string of digits and returns the number of ways it can be decoded back into its original message.

## Examples

### Example 1
**Input:** `"12"`
**Output:** `2`
**Explanation:** The possible decodings are:
- "AB" (1, 2)
- "L" (12)

### Example 2
**Input:** `"226"`
**Output:** `3`
**Explanation:** The possible decodings are:
- "BZ" (2, 26)
- "VF" (22, 6)
- "BBF" (2, 2, 6)

### Example 3
**Input:** `"0"`
**Output:** `0`
**Explanation:** There are no valid decodings (no letter maps to 0)

## Requirements

- Your solution should efficiently handle larger inputs
- Handle edge cases:
  - Leading zeros (invalid)
  - Zeros that cannot be paired (e.g., "30" is invalid)
  - Non-digit characters (invalid)
- The number of ways can grow very large for long strings

## Constraints

- The input consists of digits only
- Valid single-digit codes: 1-9
- Valid two-digit codes: 10-26
