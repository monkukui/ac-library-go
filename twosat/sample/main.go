package main

import (
	"fmt"

	"github.com/monkukui/ac-library-go/twosat"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

	/*
		AtCoder Library Practice Contest 問題H Two SAT
		https://atcoder.jp/contests/practice2/tasks/practice2_h
		例題1のサンプルコード
	*/

	var n, d int
	n, d = 3, 2

	x := make([]int, n)
	y := make([]int, n)
	x[0], y[0] = 1, 4
	x[1], y[1] = 2, 5
	x[2], y[2] = 0, 6

	ts := twosat.New(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if absInt(x[i]-x[j]) < d {
				ts.AddClause(i, false, j, false)
			}
			if absInt(x[i]-y[j]) < d {
				ts.AddClause(i, false, j, true)
			}
			if absInt(y[i]-x[j]) < d {
				ts.AddClause(i, true, j, false)
			}
			if absInt(y[i]-y[j]) < d {
				ts.AddClause(i, true, j, true)
			}
		}
	}

	if ts.Satisfiable() == false {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")

		answer := ts.Answer()
		for i := 0; i < n; i++ {
			if answer[i] == true {
				fmt.Println(x[i])
			} else {
				fmt.Println(y[i])
			}
		}
	}
}
