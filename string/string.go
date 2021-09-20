package string

import (
	"fmt"
	internal "github.com/monkukui/ac-library-go/internal/string"
	"sort"
)

const (
	ThresholdNaive    = 10
	ThresholdDoubling = 40
)

func SuffixArrayWithUpper(s []int, upper int) []int {
	ss := append([]int{}, s...)
	if upper < 0 {
		panic("upper must be non negative integer")
	}
	for _, d := range ss {
		if d < 0 || upper < d {
			panic(fmt.Sprintf("you got %d, but every element of slice must be in range of [0, %d]", d, upper))
		}
	}
	return internal.SAIS(ss, upper, ThresholdNaive, ThresholdDoubling)
}

func SuffixArrayInt(s []int) []int {
	ss := make([]int, len(s))
	copy(ss, s)
	n := len(ss)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		l, r := idx[i], idx[j]
		return ss[l] < ss[r]
	})
	s2 := make([]int, n)
	now := 0
	for i := 0; i < n; i++ {
		if i > 0 && ss[idx[i-1]] != ss[idx[i]] {
			now++
		}
		s2[idx[i]] = now
	}
	return internal.SAIS(s2, now, ThresholdNaive, ThresholdDoubling)
}

func SuffixArrayString(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return internal.SAIS(s2, 255, ThresholdNaive, ThresholdDoubling)
}

func LcpArrayInt(s, sa []int) []int {
	n := len(s)
	if n < 1 {
		panic("length of slice s must be more than or equal to 1")
	}
	rnk := make([]int, n)
	for i := 0; i < n; i++ {
		rnk[sa[i]] = i
	}
	lcp := make([]int, n-1)
	h := 0
	for i := 0; i < n; i++ {
		if h > 0 {
			h--
		}
		if rnk[i] == 0 {
			continue
		}
		j := sa[rnk[i]-1]
		for ; j+h < n && i+h < n; h++ {
			if s[j+h] != s[i+h] {
				break
			}
		}
		lcp[rnk[i]-1] = h
	}
	return lcp
}

func LcpArrayString(s string, sa []int) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return LcpArrayInt(s2, sa)
}

func ZAlgorithmInt(s []int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	z := make([]int, n)
	z[0] = 0
	j := 0
	for i := 1; i < n; i++ {
		k := &z[i]
		if j+z[j] <= i {
			*k = 0
		} else {
			*k = min(j+z[j]-i, z[i-j])
		}
		for i+*k < n && s[*k] == s[i+*k] {
			*k++
		}
	}
	z[0] = n
	return z
}

func ZAlgorithmString(s string) []int {
	n := len(s)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		s2[i] = int(s[i])
	}
	return ZAlgorithmInt(s2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
