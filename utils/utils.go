package utils

import "runtime"

// architectureBitSizeMaxSequence returns the maximum Fibonacci sequence length based on the architecture
func ArchitectureBitSizeMaxSequence() (maxFiboncciSequence int) {
	arch := runtime.GOARCH

	switch arch {
	// 64-bit
	case "amd64", "arm64", "ppc64", "ppc64le", "mips64", "mips64le", "s390x", "riscv64":
		maxFiboncciSequence = 93
	// 32-bit
	case "386", "arm", "mips", "mipsle":
		maxFiboncciSequence = 47
	// Default to 32-bit
	default:
		maxFiboncciSequence = 47
	}

	return maxFiboncciSequence
}
