package twosat

import (
	"reflect"
	"testing"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func TestTwoSat(t *testing.T) {

	/*
		AtCoder Library Practice Contest 問題H Two SAT
		https://atcoder.jp/contests/practice2/tasks/practice2_h
		例題1をテストとして実施

		input:
			n = 3
			d = 2
			X0,Y0 = 1,4
			X1,Y1 = 2,5
			X2,Y2 = 0,6
		result:
			Satisfiable() = true
			Answer() = [false, true, true]
	*/

	var n, d int
	n, d = 3, 2

	x := make([]int, n)
	y := make([]int, n)
	x[0], y[0] = 1, 4
	x[1], y[1] = 2, 5
	x[2], y[2] = 0, 6

	ts1 := New(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if absInt(x[i]-x[j]) < d {
				ts1.AddClause(i, false, j, false)
			}
			if absInt(x[i]-y[j]) < d {
				ts1.AddClause(i, false, j, true)
			}
			if absInt(y[i]-x[j]) < d {
				ts1.AddClause(i, true, j, false)
			}
			if absInt(y[i]-y[j]) < d {
				ts1.AddClause(i, true, j, true)
			}
		}
	}

	if ts1.Satisfiable() == false {
		t.FailNow()
	}

	if ans := ts1.Answer(); !reflect.DeepEqual(ans, []bool{false, true, true}) {
		t.Fatal(ans)
	}

	/*
		AtCoder Library Practice Contest 問題H Two SAT
		https://atcoder.jp/contests/practice2/tasks/practice2_h
		例題2をテストとして実施

		input:
			n = 3
			d = 3
			X0,Y0 = 1,4
			X1,Y1 = 2,5
			X2,Y2 = 0,6
		result:
			Satisfiable() = false
	*/

	d = 3
	ts2 := New(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if absInt(x[i]-x[j]) < d {
				ts2.AddClause(i, false, j, false)
			}
			if absInt(x[i]-y[j]) < d {
				ts2.AddClause(i, false, j, true)
			}
			if absInt(y[i]-x[j]) < d {
				ts2.AddClause(i, true, j, false)
			}
			if absInt(y[i]-y[j]) < d {
				ts2.AddClause(i, true, j, true)
			}
		}
	}

	if ts2.Satisfiable() == true {
		t.FailNow()
	}

}
