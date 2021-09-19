package main

import (
	"fmt"
	"github.com/monkukui/ac-library-go/segtree"
)

func main() {
	initData := []interface{}{1, 2, 3, 4, 5}
	op := func(a, b interface{}) interface{} {
		aa, _ := a.(int)
		bb, _ := b.(int)
		return aa + bb
	}
	e := func() interface{} {
		return 0
	}
	rmq := segtree.New(initData, op, e)

	fmt.Println(rmq.Get(3))
	fmt.Println(rmq.Prod(2, 4))
	rmq.Set(3, 10)
	fmt.Println(rmq.Get(3))
	fmt.Println(rmq.Prod(2, 4))
}
