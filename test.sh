#!/bin/bash
# Copyright (c) 2025 Serhii Nesterenko
# This software is released under the MIT License.

# Test script for decode-ways program
# Builds the program and runs it on test2.txt

set -e  # Exit on error

echo "Building decode-ways..."
go build -o decode-ways main.go

echo "Running on test2.txt..."
./decode-ways test2.txt
echo ""

echo "Test completed successfully!"
