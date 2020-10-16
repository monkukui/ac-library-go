package math

import (
	"math"
	"math/bits"
)

// @param m `1 <= m`
// @return x mod m
func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

// Fast moduler by barrett reduction
type Barrett struct {
	M  uint
	Im uint64
}

// @param m `1 <= m`
func NewBarrett(m uint) *Barrett {
	return &Barrett{
		M: m,
		// im: uint64(-1)/m + 1,
		Im: math.MaxUint64/uint64(m) + 1,
	}
}

// @return m
func (barrett *Barrett) Umod() uint {
	return barrett.M
}

// @param a `0 <= a < m`
// @param b `0 <= b < m`
// @return `a * b mod m`
func (barrett *Barrett) Mul(a, b uint) uint {
	z := uint64(a)
	z *= uint64(b)
	x, _ := bits.Mul64(z, barrett.Im)
	v := uint(z - x*uint64(barrett.M))
	if barrett.M <= v {
		v += barrett.M
	}
	return v
}

// @param n `0 <= n`
// @param m `1 <= m`
// @return `(x ** n) % m`
func PowMod(x, n int64, m int) int64 {
	if m == 1 {
		return 0
	}
	um := uint(m)
	r := uint64(1)
	y := uint64(SafeMod(x, int64(m)))

	for n > 0 {
		if n&1 > 0 {
			r = (r * y) % uint64(um)
		}
		y = (y * y) % uint64(um)
		n >>= 1
	}
	return int64(r)
}

// @param n `0 <= n`
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 7 || n == 61 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	d := int64(n - 1)
	for d%2 == 0 {
		d /= 2
	}
	for _, a := range []int64{2, 7, 61} {
		t := d
		y := PowMod(a, t, n)
		for t != int64(n-1) && y != 1 && y != int64(n-1) {
			y = y * y % int64(n)
			t <<= 1
		}
		if y != int64(n-1) && t%2 == 0 {
			return false
		}
	}
	return true
}

// @param b `1 <= b`
// @return (g, x) s.t. g = gcd(a, b), xa = g (mod b), 0 <= x < b/g
func InvGcd(a, b int64) (int64, int64) {
	a = SafeMod(a, b)
	if a == 0 {
		return b, 0
	}

	s := b
	t := a
	m0 := int64(0)
	m1 := int64(1)

	for t > 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u

		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}

	if m0 < 0 {
		m0 += b / s
	}
	return s, m0
}

// @param m must be prime
// @return primitive root (and minimum in now)
func PrimitiveRoot(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 {
		return 3
	}
	if m == 469762049 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	if m == 998244353 {
		return 3
	}
	var divs [20]int
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := 3; int64(i)*int64(i) <= int64(x); i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if PowMod(int64(g), int64((m-1)/divs[i]), m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}
