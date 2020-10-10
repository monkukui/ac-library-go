package math

import internal "github.com/monkukui/ac-library-go/internal/math"

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
	return r
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
