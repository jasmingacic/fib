package fibonacci

import (
	"fib/utils"
	"reflect"
	"testing"
)

func TestGenerateFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"Zero elements", 0, []int{}},
		{"Negative elements", -5, []int{}},
		{"One element", 1, []int{0}},
		{"Two elements", 2, []int{0, 1}},
		{"Five elements", 5, []int{0, 1, 1, 2, 3}},
		{"Ten elements", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{"Exceed max limit", utils.ArchitectureBitSizeMaxSequence() + 10, GenerateFibonacci(utils.ArchitectureBitSizeMaxSequence())},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateFibonacci(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generateFibonacci(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}
