package math

import (
	"fmt"
	"reflect"
	"testing"

	"math"

	"github.com/stretchr/testify/assert"
)

// https://github.com/atcoder/ac-library/blob/master/test/utils/math.hpp#L5-L17
func factors(m int) []int {
	var result []int
	for i := 2; int64(i)*int64(i) <= int64(m); i++ {
		if m%i == 0 {
			result = append(result, i)
		}
		for m%i == 0 {
			m /= i
		}
	}
	if m > 1 {
		result = append(result, m)
	}
	return result
}

// https://github.com/atcoder/ac-library/blob/master/test/utils/math.hpp#L19-L26
func isPrimitiveRoot(m, g int) bool {
	if !(1 <= g && g < m) {
		panic("")
	}
	for _, x := range factors(m - 1) {
		if PowMod(int64(g), int64((m-1)/x), m) == 1 {
			return false
		}
	}
	return true
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/utils/math_test.cpp#L6-L16
func TestFactors(t *testing.T) {
	for m := 1; m <= 500000; m++ {
		f := factors(m)
		m2 := m
		for _, x := range f {
			assert.Exactly(t, 0, m%x)
			for m2%x == 0 {
				m2 /= x
			}
		}
		assert.Exactly(t, 1, m2)
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/utils/math_test.cpp#L18-L30
func isPrimitiveRootNaive(m, g int) bool {
	if !(1 <= g && g < m) {
		panic("")
	}
	x := 1
	for i := 1; i <= m-2; i++ {
		x = int((int64)(x) * int64(g) % int64(m))
		// x == n^i
		if x == 1 {
			return false
		}
	}
	x = int(int64(x) * int64(g) % int64(m))
	if x != 1 {
		panic("")
	}
	return true
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/utils/math_test.cpp#L32-L39
func TestIsPrimitiveRootTest(t *testing.T) {
	for m := 2; m <= 500; m++ {
		if !IsPrime(m) {
			continue
		}
		for g := 1; g < m; g++ {
			assert.Exactly(t, isPrimitiveRootNaive(m, g), isPrimitiveRoot(m, g))
		}
	}
}

func gcd(a, b int64) int64 {
	if !(0 <= a && 0 <= b) {
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

// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L44-L56
func TestBarrett(t *testing.T) {
	for m := uint(1); m <= 100; m++ {
		bt := NewBarrett(m)
		for a := uint(0); a < m; a++ {
			for b := uint(0); b < m; b++ {
				assert.Exactly(t, (a*b)%m, bt.Mul(a, b))
			}
		}
	}

	bt := NewBarrett(1)
	assert.Exactly(t, uint(0), bt.Mul(0, 0))
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L58-L78
func TestBarrettBorder(t *testing.T) {
	modUpper := uint(math.MaxInt32)
	for mod := modUpper; mod >= modUpper-20; mod-- {
		bt := NewBarrett(mod)
		var v []uint
		for i := uint(0); i < 10; i++ {
			v = append(v, i)
			v = append(v, mod-i)
			v = append(v, mod/2+i)
			v = append(v, mod/2-i)
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

// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L116-L155
func TestInvGcdBound(t *testing.T) {
	var pred []int64
	for i := int64(0); i <= 10; i++ {
		pred = append(pred, i)
		pred = append(pred, -i)
		pred = append(pred, math.MinInt64+i)
		pred = append(pred, math.MaxInt64-i)

		pred = append(pred, math.MinInt64/2+i)
		pred = append(pred, math.MinInt64/2-i)
		pred = append(pred, math.MaxInt64/2+i)
		pred = append(pred, math.MaxInt64/2-i)

		pred = append(pred, math.MinInt64/3+i)
		pred = append(pred, math.MinInt64/3-i)
		pred = append(pred, math.MaxInt64/3+i)
		pred = append(pred, math.MaxInt64/3-i)

		pred = append(pred, int64(998244353))
		pred = append(pred, int64(1000000007))
		pred = append(pred, int64(1000000009))
		pred = append(pred, int64(-998244353))
		pred = append(pred, int64(-1000000007))
		pred = append(pred, int64(-1000000009))
	}

	for _, a := range pred {
		for _, b := range pred {
			if b <= 0 {
				continue
			}
			a2 := SafeMod(a, b)
			eg, x := InvGcd(a, b)
			g := gcd(a2, b)
			assert.Exactly(t, g, eg)
			assert.LessOrEqual(t, int64(0), x)
			assert.LessOrEqual(t, x, b/eg)
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L157-L172
func TestPrimitiveRootNaive(t *testing.T) {
	for m := 2; m <= 1000; m++ {
		if !IsPrime(m) {
			continue
		}
		n := PrimitiveRoot(m)
		assert.LessOrEqual(t, 1, n)
		assert.Less(t, n, m)
		x := 1
		for i := 1; i <= m-2; i++ {
			x = int(int64(x) * int64(n) % int64(m))
			assert.NotEqual(t, 1, x)
		}
		x = int(int64(x) * int64(n) % int64(m))
		assert.Exactly(t, 1, x)
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/internal_math_test.cpp#L199-L206
func TestPrimitiveRoot(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x := math.MaxInt32 - i
		if !IsPrime(x) {
			continue
		}
		assert.True(t, isPrimitiveRoot(x, PrimitiveRoot(x)))
	}
}
