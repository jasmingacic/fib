package utils

import (
	"runtime"
	"testing"
)

func TestArchitectureBitSizeMaxSequence(t *testing.T) {
	expected := map[string]int{
		"amd64":    93,
		"arm64":    93,
		"ppc64":    93,
		"ppc64le":  93,
		"mips64":   93,
		"mips64le": 93,
		"s390x":    93,
		"riscv64":  93,
		"386":      47,
		"arm":      47,
		"mips":     47,
		"mipsle":   47,
	}

	arch := runtime.GOARCH
	got := ArchitectureBitSizeMaxSequence()

	if want, ok := expected[arch]; ok {
		if got != want {
			t.Errorf("For architecture %s, expected %d but got %d", arch, want, got)
		}
	} else {
		if got <= 0 {
			t.Errorf("Unexpected architecture %s returned invalid sequence length %d", arch, got)
		}
	}
}
