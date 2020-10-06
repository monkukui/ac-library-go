package fenwicktree

type Int struct {
	n    int
	data []uint
}

func NewInt(n int) *Int {
	i := &Int{
		n:    n,
		data: make([]uint, n),
	}
	for idx := range i.data {
		i.data[idx] = 0
	}
	return i
}

func (i *Int) Add(pos int, x int) {
	if !(0 <= pos && pos < i.n) {
		panic("")
	}
	pos++
	for pos <= i.n {
		i.data[pos-1] += uint(x)
		pos += pos & -pos
	}
}

func (i *Int) Sum(l, r int) int {
	if !(0 <= l && l <= r && r <= i.n) {
		panic("")
	}
	return int(i.sum(r) - i.sum(l))
}

func (i *Int) sum(r int) uint {
	s := uint(0)
	for r > 0 {
		s += i.data[r-1]
		r -= r & -r
	}
	return s
}
