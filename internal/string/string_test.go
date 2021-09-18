package string

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L13-L22
func saNaive(s []int) []int {
	n := len(s)
	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	sort.Slice(sa, func(l, r int) bool {
		for ; l < n && r < n; {
			if s[l] != s[r] {
				return s[l] < s[r]
			}
			l++
			r++
		}
		return l == n
	})
	return sa
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L24-L33
func lcpNaive(s, sa []int) []int {
	n := len(s)
	if n == 0 {
		panic("")
	}
	lcp := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		l, r := sa[i], sa[i+1]
		for l+lcp[i] < n && r+lcp[i] < n && s[l+lcp[i]] == s[r+lcp[i]] {
			lcp[i]++
		}
	}
	return lcp
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L91-L124
func TestInternalSANaive(t *testing.T) {
	for n := 1; n <= 5; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 4
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			for i := 0; i < n; i++ {
				s[i] = g % 4
				g /= 4
			}
			assert.Equal(t, saNaive(s), SANaive(s))
		}
	}
	for n := 1; n <= 10; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 2
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			for i := 0; i < n; i++ {
				s[i] = g % 2
				g /= 2
			}
			assert.Equal(t, saNaive(s), SANaive(s))
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L126-L157
func TestInternalSADoubling(t *testing.T) {
	for n := 1; n <= 5; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 4
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			for i := 0; i < n; i++ {
				s[i] = g % 4
				g /= 4
			}
			assert.Equal(t, saNaive(s), SADoubling(s))
		}
	}
	for n := 1; n <= 10; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 2
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			for i := 0; i < n; i++ {
				s[i] = g % 2
				g /= 2
			}
			assert.Equal(t, saNaive(s), SADoubling(s))
		}
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L159-L194
func TestInternalSAIS(t *testing.T) {
	for n := 1; n <= 5; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 4
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			maxc := 0
			for i := 0; i < n; i++ {
				s[i] = g % 4
				if maxc < s[i] {
					maxc = s[i]
				}
				g /= 4
			}
			assert.Equal(t, saNaive(s), SAIS(s, maxc, -1, -1))
		}
	}
	for n := 1; n <= 10; n++ {
		m := 1
		for i := 0; i < n; i++ {
			m *= 2
		}
		for f := 0; f < m; f++ {
			s := make([]int, n)
			g := f
			maxc := 0
			for i := 0; i < n; i++ {
				s[i] = g % 2
				if maxc < s[i] {
					maxc = s[i]
				}
				g /= 2
			}
			assert.Equal(t, saNaive(s), SAIS(s, maxc, -1, -1))
		}
	}
}

