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
	fmt.Println("set d[", 1, "] = ", 20)
	rmq.Set(1, 20)

	fmt.Println("d[", 1, "] = ", rmq.Get(1))
	fmt.Println("d[", 3, "] = ", rmq.Get(3))
	fmt.Println("Prod(", 0, ",", 4, ") = ", rmq.Prod(0, 4))
	fmt.Println(rmq.Prod(0, 4))

	fmt.Println("set d[", 3, "] = ", 10)
	rmq.Set(3, 10)

	fmt.Println("d[", 1, "] = ", rmq.Get(1))
	fmt.Println("d[", 3, "] = ", rmq.Get(3))
	fmt.Println("Prod(", 0, ",", 2, ") = ", rmq.Prod(0, 2))
	fmt.Println("Prod(", 0, ",", 4, ") = ", rmq.Prod(0, 4))
}
