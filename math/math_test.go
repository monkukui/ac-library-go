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

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L61-L68
func TestInvBoundHand(t *testing.T) {
	minll := int64(math.MinInt64)
	maxll := int64(math.MaxInt64)
	assert.Exactly(t, InvMod(-1, maxll), InvMod(minll, maxll))
	assert.Exactly(t, int64(1), InvMod(maxll, maxll-1))
	assert.Exactly(t, maxll-1, InvMod(maxll-1, maxll))
	assert.Exactly(t, int64(2), InvMod(maxll/2+1, maxll))
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L70-L80
func TestInvMod(t *testing.T) {
	for a := int64(-100); a <= 100; a++ {
		for b := int64(1); b <= 1000; b++ {
			if gcd(internal.SafeMod(a, b), b) != 1 {
				continue
			}
			c := InvMod(a, b)
			assert.LessOrEqual(t, int64(0), c)
			assert.Less(t, c, b)
			assert.Exactly(t, 1%b, ((a*c)%b+b)%b)
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L82-L90
func TestInvModZero(t *testing.T) {
	assert.Exactly(t, int64(0), InvMod(0, 1))
	for i := int64(0); i < 10; i++ {
		assert.Exactly(t, int64(0), InvMod(i, int64(1)))
		assert.Exactly(t, int64(0), InvMod(-i, int64(1)))
		assert.Exactly(t, int64(0), InvMod(math.MinInt64+i, int64(1)))
		assert.Exactly(t, int64(0), InvMod(math.MaxInt64-i, int64(1)))
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L92-L103
func TestFloorSum(t *testing.T) {
	for n := int64(0); n < 20; n++ {
		for m := int64(1); m < 20; m++ {
			for a := int64(0); a < 20; a++ {
				for b := int64(0); b < 20; b++ {
					assert.Exactly(t, floorSumNaive(n, m, a, b), FloorSum(n, m, a, b))
				}
			}
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L105-L109
func TestCrtHand(t *testing.T) {
	rem, mod := Crt([]int64{1, 2, 1}, []int64{2, 3, 2})
	assert.Exactly(t, int64(5), rem)
	assert.Exactly(t, int64(6), mod)
}
