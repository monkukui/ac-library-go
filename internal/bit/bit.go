// https://github.com/atcoder/ac-library/blob/master/atcoder/internal_bit.hpp
package bit

import (
	"math/bits"
)

// @param n `0 <= n`
// @return minimum non-negative `x` s.t. `n <= 2**x`
func CeilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}

// @param n `1 <= n`
// @return minimum non-negative `x` s.t. `(n & (1 << x)) != 0`
func Bsf(n uint) int {
	return bits.TrailingZeros(n)
}
