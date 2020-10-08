package math

import (
	"testing"

	"math"

	"github.com/stretchr/testify/assert"
)

func gcd(a, b int64) int64 {
	if 0 <= a && a <= b {
		panic("")
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func powModNaive(x, n int64, mod int) int64 {
	y := (int(x%int64(mod)) + mod) % mod
	z := 1
	for i := 0; int64(i) < n; i++ {
		z = (z * y) % mod
	}
	return int64(z % mod)
}

func floorSumNaive(n, m, a, b int64) int64 {
	sum := int64(0)
	for i := int64(0); i < n; i++ {
		sum += (a*i + b) / m
	}
	return sum
}

func isPrimeNaive(n int) bool {
	if 0 <= n && n <= math.MaxInt32 {
		panic("")
	}
	if n == 0 || n == 1 {
		return false
	}
	for i := int64(2); i*i <= int64(n); i++ {
		if n%int(i) == 0 {
			return false
		}
	}
	return true
}

// Test Barrett は、mod m での乗算が正しく行えることをテストします。
func TestBarrett(t *testing.T) {
	for m := uint(1); m <= 100; m++ {
		bt := New(m)
		for a := uint(0); a < m; a++ {
			for b := uint(0); b < m; b++ {
				assert.Exactly(t, (a*b)%m, bt.Mul(a, b))
			}
		}
	}

	bt := New(1)
	assert.Exactly(t, uint(0), bt.Mul(0, 0))
}
