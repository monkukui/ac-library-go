package string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		for  i := 0; i < n; i++ {
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
		for i + z[i] < n && s[z[i]] == s[i + z[i]] {
			z[i]++
		}
	}
	return z
}
