package math

import (
	"math"
	"testing"

	internal "github.com/monkukui/ac-library-go/internal/math"

	"github.com/stretchr/testify/assert"
)

func gcd(a, b int64) int64 {
	if !(0 <= a && 0 <= b) {
		panic("")
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// TODO 確認：この鬼のようなキャストはなんとかならないか？
func powModNaive(x int64, n uint64, mod uint) int64 {
	y := uint64(uint64(x)%uint64(mod)+uint64(mod)) % uint64(mod)
	z := uint64(1)
	for i := uint64(0); i < n; i++ {
		z = (z * uint64(y)) % uint64(mod)
	}
	return int64(z % uint64(mod))
}

func floorSumNaive(n, m, a, b int64) int64 {
	sum := int64(0)
	for i := int64(0); i < n; i++ {
		sum += (a*i + b) / m
	}
	return sum
}

func isPrimeNaive(n int64) bool {
	if !(0 <= n && n <= math.MaxInt32) {
		panic("")
	}
	if n == 0 || n == 1 {
		return false
	}
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L43-L59
func TestPowMod(t *testing.T) {
	naive := func(x, n int64, mod int) int64 {
		y := internal.SafeMod(x, int64(mod))
		z := uint64(1) % uint64(mod)
		for i := int64(0); i < n; i++ {
			z = (z * uint64(y)) % uint64(mod)
		}
		return int64(z)
	}
	for a := int64(-100); a <= 100; a++ {
		for b := int64(0); b <= 100; b++ {
			for c := 1; c <= 100; c++ {
				assert.Exactly(t, naive(a, b, c), PowMod(a, b, c))
			}
		}
	}
}
