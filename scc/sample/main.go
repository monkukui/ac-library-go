package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/monkukui/ac-library-go/scc"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// AC code of https://atcoder.jp/contests/practice2/tasks/practice2_g
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	n, m := getInt(), getInt()

	g := scc.NewGraph(n)
	for i := 0; i < m; i++ {
		u, v := getInt(), getInt()
		g.AddEdge(u, v)
	}

	scc := g.Scc()

	out(len(scc))
	for _, e := range scc {
		fmt.Print(len(e))
		for _, x := range e {
			fmt.Print(" ", x)
		}
		out()
	}
}
