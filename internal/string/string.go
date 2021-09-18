package string

import "sort"

func saNaive(s []int) []int {
	n := len(s)
	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	sort.Slice(sa, func(l, r int) bool {
		if l == r {
			return false
		}
		for l < n && r < n {
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

func saDoubling(s []int) []int {
	n := len(s)
	sa, rnk, tmp := make([]int, n), s, make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	for k := 1; k < n; k *= 2 {
		cmp := func(x, y int) bool {
			if rnk[x] != rnk[y] {
				return rnk[x] < rnk[y]
			}
			rx, ry := -1, -1
			if x+k < n {
				rx = rnk[x+k]
			}
			if y+k < n {
				ry = rnk[y+k]
			}
			return rx < ry
		}
		sort.Slice(sa, cmp)
		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			if cmp(sa[i-1], sa[i]) {
				tmp[sa[i]]++
			}
		}
		tmp, rnk = rnk, tmp
	}
	return sa
}

const (
	THRESHOLD_NAIVE    = 10
	THRESHOLD_DOUBLING = 40
)

func sais(s []int, upper int) []int {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if s[0] < s[1] {
			return []int{0, 1}
		} else {
			return []int{1, 0}
		}
	}
	if n < THRESHOLD_NAIVE {
		return saNaive(s)
	}
	if n < THRESHOLD_DOUBLING {
		return saDoubling(s)
	}

	sa := make([]int, n)
	ls := make([]bool, n)
	for i := n - 2; i >= 0; i-- {
		if s[i] == s[i+1] {
			ls[i] = ls[i+1]
		} else {
			ls[i] = s[i] < s[i+1]
		}
	}
	suml, sums := make([]int, upper+1), make([]int, upper+1)
	for i := 0; i < n; i++ {
		if !ls[i] {
			sums[s[i]]++
		} else {
			suml[s[i]+1]++
		}
	}
	for i := 0; i <= upper; i++ {
		sums[i] += suml[i]
		if i < upper {
			suml[i+1] += sums[i]
		}
	}

	induce := func(lms []int) {
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		buf := make([]int, upper+1)
		copy(buf, sums)
		for _, d := range lms {
			if d == n {
				continue
			}
			sa[buf[s[d]]] = d
			buf[s[d]]++
		}
		copy(buf, suml)
		sa[buf[s[n-1]]] = n - 1
		buf[s[n-1]]++
		for i := 0; i < n; i++ {
			v := sa[i]
			if v >= 1 && !ls[v-1] {
				sa[buf[s[v-1]]] = v - 1
				buf[s[v-1]]++
			}
		}
		copy(buf, suml)
		for i := n - 1; i >= 0; i-- {
			v := sa[i]
			if v >= 1 && ls[v - 1] {
				sa[buf[s[v - 1] + 1] - 1] = v - 1
				buf[s[v - 1] + 1]--
			}
		}
	}
	lmsMap := make([]int, n + 1)
	for i := 0; i < len(lmsMap); i++ {
		lmsMap[i] = -1
	}
	m := 0
	for i := 1; i < n; i++ {
		if !ls[i - 1] && ls[i] {
			lmsMap[i] = m
			m++
		}
	}
	lms := make([]int, 0, m)
	for i := 1; i < n; i++ {
		if !ls[i - 1] && ls[i] {
			lms = append(lms, i)
		}
	}
	induce(lms)
	if m > 0 {
		sortedLms := make([]int, 0, m)
		for _, v := range(sa) {
			if lmsMap[v] != -1 {
				sortedLms = append(sortedLms, v)
			}
		}
		recs := make([]int, m)
		recUpper := 0
		recs[lmsMap[sortedLms[0]]] = 0
		for i := 1; i < m; i++ {
			l, r := sortedLms[i - 1], sortedLms[i]
			endl, endr := n, n
			if lmsMap[l] + 1 < m {
				endl = lms[lmsMap[l] + 1]
			}
			if lmsMap[r] + 1 < m {
				endr = lms[lmsMap[r] + 1]
			}
			same := true
			if endl - l != endr - r {
				same = false
			} else {
				for l < endl {
					if s[l] != s[r] {
						break
					}
					l++
					r++
				}
				if l == n || s[l] != s[r] {
					same = false
				}
			}
			if !same {
				recUpper++
			}
			recs[lmsMap[sortedLms[i]]] = recUpper
		}
		recSa := sais(recs, recUpper)
		for i := 0; i < m; i++ {
			sortedLms[i] = lms[recSa[i]]
		}
		induce(sortedLms)
	}
	return sa
}
