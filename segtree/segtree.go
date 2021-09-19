package segtree

type SegTree struct {
	d []interface{}
	op func(a, b interface{}) interface{}
	e func() interface{}
}

func New(d []interface{}, op func(a, b interface{}) interface{}, e func() interface{}) *SegTree {
	return &SegTree{
		d: d,
		op: op,
		e: e,
	}
}

func (s *SegTree) Set(p int, x interface{}) {
	s.d[p] = x
}

func (s *SegTree) Get(p int) interface{} {
	return s.d[p]
}

func (s *SegTree) Prod(l, r int) interface{} {
	v := s.e()
	for i := l; i < r; i++ {
		v = s.op(v, s.d[i])
	}
	return v
}

func (s *SegTree) AllProd() interface{} {
	return s.Prod(0, len(s.d))
}
