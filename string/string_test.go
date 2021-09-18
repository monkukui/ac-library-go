package string

import (
	internal "github.com/monkukui/ac-library-go/internal/string"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

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

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L196-L203
func TestSAAllA(t *testing.T) {
	for n := 1; n <= 100; n++ {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = 10
		}
		assert.Equal(t, internal.SANaive(s), SuffixArrayInt(s))
		assert.Equal(t, internal.SANaive(s), SuffixArrayWithUpper(s, 10))
		assert.Equal(t, internal.SANaive(s), SuffixArrayWithUpper(s, 12))
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L205-L218
func TestSAAllAB(t *testing.T) {
	for n := 1; n <= 100; n++ {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i % 2
		}
		assert.Equal(t, internal.SANaive(s), SuffixArrayInt(s))
		assert.Equal(t, internal.SANaive(s), SuffixArrayWithUpper(s, 3))
	}
	for n := 1; n <= 100; n++ {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = 1 - i%2
		}
		assert.Equal(t, internal.SANaive(s), SuffixArrayInt(s))
		assert.Equal(t, internal.SANaive(s), SuffixArrayWithUpper(s, 3))
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L220-L243
func TestSAString(t *testing.T) {
	s := "missisippi"
	sa := SuffixArrayString(s)
	answer := []string{
		"i",          // 9
		"ippi",       // 6
		"isippi",     // 4
		"issisippi",  // 1
		"missisippi", // 0
		"pi",         // 8
		"ppi",        // 7
		"sippi",      // 5
		"sisippi",    // 3
		"ssisippi",   // 2
	}
	assert.Equal(t, len(answer), len(sa))
	for i := 0; i < len(sa); i++ {
		assert.Equal(t, answer[i], s[sa[i]:])
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L245-L253
func TestSASingle(t *testing.T) {
	assert.Equal(t, []int{0}, SuffixArrayInt([]int{0}))
	assert.Equal(t, []int{0}, SuffixArrayInt([]int{-1}))
	assert.Equal(t, []int{0}, SuffixArrayInt([]int{1}))
	assert.Equal(t, []int{0}, SuffixArrayInt([]int{math.MinInt32}))
	assert.Equal(t, []int{0}, SuffixArrayInt([]int{math.MaxInt32}))
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L255-L286
func TestLCP(t *testing.T) {
	s := "aab"
	sa := SuffixArrayString(s)
	assert.Equal(t, []int{0, 1, 2}, sa)
	lcp := LcpArrayString(s, sa)
	assert.Equal(t, []int{1, 0}, lcp)

	assert.Equal(t, lcp, LcpArrayInt([]int{0, 0, 1}, sa))
	assert.Equal(t, lcp, LcpArrayInt([]int{-100, -100, 100}, sa))
	assert.Equal(t, lcp, LcpArrayInt([]int{
		math.MinInt32,
		math.MinInt32,
		math.MaxInt32,
	}, sa))
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L52-L89
func TestSALCP(t *testing.T) {
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
			sa := internal.SANaive(s)
			assert.Equal(t, sa, SuffixArrayInt(s))
			assert.Equal(t, sa, SuffixArrayWithUpper(s, maxc))
			assert.Equal(t, lcpNaive(s, sa), LcpArrayInt(s, sa))
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
			sa := internal.SANaive(s)
			assert.Equal(t, sa, SuffixArrayInt(s))
			assert.Equal(t, sa, SuffixArrayWithUpper(s, maxc))
			assert.Equal(t, lcpNaive(s, sa), LcpArrayInt(s, sa))
		}
	}
}

func TestZAlgorithmString_HandMaid(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		data     string
		expected []int
	}{
		{
			name:     "引数が空文字列の時に期待通り動作する",
			data:     "",
			expected: []int{},
		},
		{
			name:     "引数が abab の時に期待通り動作する",
			data:     "abab",
			expected: []int{4, 0, 2, 0},
		},
		{
			name:     "引数が aaaaa の時に期待通り動作する",
			data:     "aaaaa",
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "引数が aaabaaaab の時に期待通り動作する", // https://snuke.hatenablog.com/entry/2014/12/03/214243
			data:     "aaabaaaab",
			expected: []int{9, 2, 1, 0, 3, 4, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, ZAlgorithmString(tt.data))
		})
	}
}
func TestZAlgorithmInt_HandMaid(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		name     string
		data     []int
		expected []int
	}{
		{
			name:     "引数が空スライスの時に期待通り動作する",
			data:     []int{},
			expected: []int{},
		},
		{
			name:     "引数が {1, 2, 1, 2} の時に期待通り動作する",
			data:     []int{1, 2, 1, 2},
			expected: []int{4, 0, 2, 0},
		},
		{
			name:     "引数が {1, 1, 1, 1, 1} の時に期待通り動作する",
			data:     []int{1, 1, 1, 1, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "引数が {1, 1, 1, 2, 1, 1, 1, 1, 2} の時に期待通り動作する",
			data:     []int{1, 1, 1, 2, 1, 1, 1, 1, 2},
			expected: []int{9, 2, 1, 0, 3, 4, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expected, ZAlgorithmInt(tt.data))
		})
	}
}

// https://github.com/atcoder/ac-library/blob/master/test/unittest/string_test.cpp#L297-L324
func TestZAlgorithmInt_CompareWithNaive(t *testing.T) {
	t.Parallel()
	for n := 1; n <= 6; n++ {
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
			assert.Equal(t, ZNaiveInt(s), ZAlgorithmInt(s))
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
			assert.Equal(t, ZNaiveInt(s), ZAlgorithmInt(s))
		}
	}
}

func ZNaiveInt(s []int) []int {
	n := len(s)
	z := make([]int, n)
	for i := 0; i < n; i++ {
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
	}
	return z
}
