package main

import (
	"fmt"
	"github.com/monkukui/ac-library-go/segtree"
	"math"
)

func main() {
	op := func(a, b interface{}) interface{} {
		aa, _ := a.(int)
		bb, _ := b.(int)
		if aa < bb {
			return aa
		}
		return bb
	}
	e := func() interface{} {
		return math.MaxInt64
	}
	rmq := segtree.New(op, e, 10)

	fmt.Println(rmq.Get(3))
	fmt.Println(rmq.Prod(2, 4))
	rmq.Set(3, 10)
	fmt.Println(rmq.Get(3))
	fmt.Println(rmq.Prod(2, 4))
}
