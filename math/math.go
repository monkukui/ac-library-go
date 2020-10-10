package math

import internal "github.com/monkukui/ac-library-go/internal/math"

func swap(a, b int64) (int64, int64) {
	return b, a
}

func PowMod(x, n int64, m int) int64 {
	if !(0 <= n && 1 <= m) {
		panic("")
	}
	if m == 1 {
		return 0
	}
	bt := internal.New(uint(m))
	r := uint(1)
	y := uint(internal.SafeMod(x, int64(m)))
	for n > 0 {
		if n&1 > 0 {
			r = bt.Mul(r, y)
		}
		y = bt.Mul(y, y)
		n >>= 1
	}
	return int64(r)
}

func InvMod(x, m int64) int64 {
	if !(1 <= m) {
		panic("")
	}
	g, x := internal.InvGcd(x, m)
	if g != 1 {
		panic("")
	}
	return x
}

// (rem, mod)
func Crt(r, m []int64) (int64, int64) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0 := int64(0)
	m0 := int64(1)
	for i := 0; i < n; i++ {
		if !(1 <= m[i]) {
			panic("")
		}
		r1 := internal.SafeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = swap(r0, r1)
			m0, m1 = swap(m0, m1)
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := internal.InvGcd(m0, m1)

		u1 := m1 / g
		if (r1-r0)%g > 1 {
			return 0, 0
		}

		x := (r1 - r0) / g % u1 * im % u1

		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}

func FloorSum(n, m, a, b int64) int64 {
	ans := int64(0)
	if a >= m {
		ans += (n - 1) * n * (a / m) / 2
		a %= m
	}
	if b >= m {
		ans += n * (b / m)
		b %= m
	}

	yMax := (a*n + b) / m
	xMax := (yMax*m - b)
	if yMax == 0 {
		return ans
	}
	ans += (n - (xMax+a-1)/a) * yMax
	ans += FloorSum(yMax, a, m, (a-xMax%a)%a)
	return ans
}
