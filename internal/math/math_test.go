package math

import (
	"fmt"
	"reflect"
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
	if !(0 <= n && n <= math.MaxInt32) {
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

// TestBarrett は、mod m での乗算が正しく行えることをテストします。
// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L44-L56
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

// TestBarrettBorder は、mod m での乗算が、境界でも正しく行えることをテストします。
// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L58-L78
func TestBarrettBorder(t *testing.T) {
	modUpper := uint(math.MaxInt32)
	for mod := modUpper; mod >= modUpper-20; mod-- {
		bt := New(mod)
		v := make([]uint, 10*4)
		for i := uint(0); i < 10; i++ {
			v[i*4+0] = i
			v[i*4+1] = mod - i
			v[i*4+2] = mod/2 + i
			v[i*4+3] = mod/2 - i
		}
		for _, a := range v {
			a2 := int64(a)
			if a2*a2%int64(mod)*a2%int64(mod) < 0 {
				fmt.Println(a2, a2*a2, reflect.TypeOf(a2), reflect.TypeOf(a))
				fmt.Println(math.MaxUint32, math.MaxInt64)
				fmt.Println(bt.Mul(a, bt.Mul(a, a)))
			}
			assert.Exactly(t, uint(((a2*a2)%int64(mod)*a2)%int64(mod)), bt.Mul(a, bt.Mul(a, a)))
			for _, b := range v {
				b2 := int64(b)
				assert.Equal(t, uint((a2*b2)%int64(mod)), bt.Mul(a, b))
			}
		}
	}
}

// TestIsPrime は、素数判定が正しく行えることをテストします。
// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L80-L93
func TestIsPrime(t *testing.T) {
	assert.False(t, IsPrime(121))
	assert.False(t, IsPrime(11*13))
	assert.True(t, IsPrime(1000000007))
	assert.False(t, IsPrime(1000000008))
	assert.True(t, IsPrime(1000000009))

	for i := 0; i <= 10000; i++ {
		assert.Exactly(t, isPrimeNaive(i), IsPrime(i))
	}

	for i := 0; i <= 10000; i++ {
		x := math.MaxInt32 - i
		assert.Exactly(t, isPrimeNaive(x), IsPrime(x))
	}
}
