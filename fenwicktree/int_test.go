package fenwicktree

import (
	"math/rand"
	"testing"
)

func TestInt(t *testing.T) {
	bit := NewInt(5)
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
