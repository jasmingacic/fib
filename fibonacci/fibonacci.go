package fibonacci

import "fib/utils"

// generateFibonacci returns the first n Fibonacci numbers starting from 0
// time and space complexity is O(n)
func GenerateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	// check if n is greater than the maximum Fibonacci sequence length for the current architecture if exceeeds then return max
	if n > utils.ArchitectureBitSizeMaxSequence() {
		n = utils.ArchitectureBitSizeMaxSequence()
	}

	fib := make([]int, n)

	// Initialize the first two Fibonacci numbers
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fib[i] = a
		a, b = b, a+b
	}

	return fib
}

// generateFibonacci returns the first n Fibonacci numbers starting from 0
// This solution is easier to understand and maintain than the previous one.
// complexity is O(n) and space complexity is O(n) just like the previous one
func generateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	// check if n is greater than the maximum Fibonacci sequence length for the current architecture if exceeeds then return max
	if n > utils.ArchitectureBitSizeMaxSequence() {
		n = utils.ArchitectureBitSizeMaxSequence()
	}

	fib := make([]int, n)
	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
