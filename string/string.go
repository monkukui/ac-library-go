package string

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
		if j + z[j] <= i {
			*k = 0
		} else {
			*k = min(j + z[j] - i, z[i - j])
		}
		for i + *k < n && s[*k] == s[i + *k] {
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
