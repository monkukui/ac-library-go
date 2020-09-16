package fenwicktree

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestInt(t *testing.T) {
	bit := NewInt(5)

	// check if all values are zero
	for i := 0; i < 5; i++ {
		if bit.Sum(i, i+1) != 0 {
			t.Fatal(i, bit.Sum(i, i+1))
		}
	}

	// [1, 2, 3, 4, 5]
	for i := 0; i < 5; i++ {
		bit.Add(i, i+1)
	}
	if bit.Sum(0, 5) != 15 {
		t.FailNow()
	}
	if bit.Sum(0, 4) != 10 {
		t.FailNow()
	}
	if bit.Sum(1, 3) != 5 {
		t.FailNow()
	}
}

// TestIntSumOverflow
// (*Int).Sum returns the answer in $\bmod 2^{\mathrm{bit}}$, if overflowed.
func TestIntSumOverflow(t *testing.T) {
	const N = 5000
	const intSize = 32 << (^uint(0) >> 63)

	// mod := 1 << intSize
	mod := big.NewInt(0).SetUint64(uint64((1 << intSize) - 1))
	mod.Add(mod, big.NewInt(1))

	// returns if x ≡ y (mod 2^intSize)
	congruence := func(x, y *big.Int) bool {
		xmod := big.NewInt(0).Mod(x, mod)
		ymod := big.NewInt(0).Mod(y, mod)
		return xmod.Cmp(ymod) == 0
	}

	bit := NewInt(N)
	sum := big.NewInt(0)

	for i := 0; i < N; i++ {
		r := rand.Int()
		bit.Add(i, r)
		sum.Add(sum, big.NewInt(int64(r)))
		if !congruence(sum, big.NewInt(int64(bit.Sum(0, i+1)))) {
			t.FailNow()
		}
	}
}

func BenchmarkInt(b *testing.B) {
	// https://atcoder.jp/contests/practice2/tasks/practice2_b

	// generate testcase
	const N = 500000
	Q := b.N
	const MaxInput = 1000000000

	a := make([]int, N)
	queries := make([][]int, Q)

	for i := range a {
		a[i] = int(rand.Int31n(MaxInput))
	}
	for i := range queries {
		q := []int{int(rand.Int31n(2))}
		switch q[0] {
		case 0:
			q = append(q, []int{int(rand.Int31n(N)), int(rand.Int31n(MaxInput))}...)
		case 1:
			q = append(q, int(rand.Int31n(N)))
			// (q[1], N]
			q = append(q, q[1]+1+int(rand.Int31n(int32(N-q[1]))))
		}
		queries[i] = q
	}

	bit := NewInt(N)
	for i, v := range a {
		bit.Add(i, v)
	}
	b.ResetTimer()
	for _, q := range queries {
		switch q[0] {
		case 0:
			p, x := q[1], q[2]
			bit.Add(p, x)
		case 1:
			l, r := q[1], q[2]
			bit.Sum(l, r)
		}
	}
}
