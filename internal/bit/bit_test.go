// https://github.com/atcoder/ac-library/blob/master/atcoder/internal_bit.hpp
package bit

import (
	"math"
	"testing"
)

func TestCeilPow2(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		// https://github.com/atcoder/ac-library/blob/2088c8e2431c3f4d29a2cfabc6529fe0a0586c48/test/unittest/bit_test.cpp
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 4},
		{1 << 30, 30},
		{(1 << 30) + 1, 31},
		{math.MaxInt32, 31},
	}
	for _, tt := range tests {
		if got := CeilPow2(tt.n); got != tt.want {
			t.Errorf("CeilPow2(%v) = %d, want %v", tt.n, got, tt.want)
		}
	}
}
