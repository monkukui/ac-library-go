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

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L111-L130
func TestCrt2(t *testing.T) {
	for a := int64(1); a <= 20; a++ {
		for b := int64(1); b <= 20; b++ {
			for c := int64(-10); c <= 10; c++ {
				for d := int64(-10); d <= 10; d++ {
					rem, mod := Crt([]int64{c, d}, []int64{a, b})
					if mod == 0 {
						for x := int64(0); x < a*b/gcd(a, b); x++ {
							assert.True(t, x%a != c || x%b != d)
						}
						continue
					}
					assert.Exactly(t, a*b/gcd(a, b), mod)
					assert.Exactly(t, internal.SafeMod(c, a), rem%a)
					assert.Exactly(t, internal.SafeMod(d, b), rem%b)
				}
			}
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L132-L159
func TestCrt3(t *testing.T) {
	for a := int64(1); a <= 5; a++ {
		for b := int64(1); b <= 5; b++ {
			for c := int64(1); c <= 5; c++ {
				for d := int64(-5); d <= 5; d++ {
					for e := int64(-5); e <= 5; e++ {
						for f := int64(-5); f <= 5; f++ {
							rem, mod := Crt([]int64{d, e, f}, []int64{a, b, c})
							lcm := a * b / gcd(a, b)
							lcm = lcm * c / gcd(lcm, c)
							if mod == 0 {
								for x := int64(0); x < lcm; x++ {
									assert.True(t, x%a != d || x%b != e || x%c != f)
								}
								continue
							}

							assert.Exactly(t, lcm, mod)
							assert.Exactly(t, internal.SafeMod(d, a), rem%a)
							assert.Exactly(t, internal.SafeMod(e, b), rem%b)
							assert.Exactly(t, internal.SafeMod(f, c), rem%c)
						}
					}
				}
			}
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L161-L170
func TestCrtOverflow(t *testing.T) {
	r0 := int64(0)
	r1 := int64(1000000000000 - 2)
	m0 := int64(900577)
	m1 := int64(1000000000000)
	rem, mod := Crt([]int64{r0, r1}, []int64{m0, m1})
	assert.Exactly(t, m0*m1, mod)
	assert.Exactly(t, r0, rem%m0)
	assert.Exactly(t, r1, rem%m1)
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/math_test.cpp#L172-L227
func TestCrtBound(t *testing.T) {
	INF := int64(math.MaxInt64)
	var pred []int64
	for i := int64(1); i <= 10; i++ {
		pred = append(pred, i)
		pred = append(pred, INF-(i-1))
	}
	pred = append(pred, 998244353)
	pred = append(pred, 1000000007)
	pred = append(pred, 1000000009)

	for _, ab := range [][2]int64{
		{INF, INF},
		{1, INF},
		{INF, 1},
		{7, INF},
		{INF / 337, 337},
		{2, (INF - 1) / 2}} {

		a := ab[0]
		b := ab[1]
		for ph := 0; ph < 2; ph++ {
			for _, ans := range pred {
				rem, mod := Crt([]int64{ans % a, ans % b}, []int64{a, b})
				lcm := a / gcd(a, b) * b
				assert.Exactly(t, lcm, mod)
				assert.Exactly(t, ans%lcm, rem)
			}
			a, b = swap(a, b)
		}
	}

	for _, factorInf := range permutations([]int64{49, 73, 127, 337, 92737, 649657}) {
		for _, ans := range pred {
			var r, m []int64
			for _, f := range factorInf {
				r = append(r, ans%f)
				m = append(m, f)
			}
			rem, mod := Crt(r, m)
			assert.Exactly(t, ans%INF, rem)
			assert.Exactly(t, INF, mod)
		}
	}

	for _, factorInf := range permutations([]int64{2, 3, 715827883, 2147483647}) {
		for _, ans := range pred {
			var r, m []int64
			for _, f := range factorInf {
				r = append(r, ans%f)
				m = append(m, f)
			}
			rem, mod := Crt(r, m)
			assert.Exactly(t, ans%(INF-1), rem)
			assert.Exactly(t, INF-1, mod)
		}
	}
}

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func permutations(arr []int64) [][]int64 {
	var helper func([]int64, int)
	res := [][]int64{}

	helper = func(arr []int64, n int) {
		if n == 1 {
			tmp := make([]int64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func TestPermutations(t *testing.T) {
	arr := []int64{1, 2, 3}
	assert.Exactly(t, permutations(arr), [][]int64{
		{1, 2, 3},
		{2, 1, 3},
		{3, 2, 1},
		{2, 3, 1},
		{3, 1, 2},
		{1, 3, 2}})
}
