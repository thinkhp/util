package thinkMath

import "sort"

func HaveIntersection(n1, n2 []int) []int{
	sort.Ints(n1)
	sort.Ints(n2)

	// find
	i := 0
	j := 0
	result := make([]int, 0)
	for {
		if i >= len(n1) || j >= len(n2) {
			break
		}
		if n1[i] == n2[j] {
			result = append(result, n1[i])
			i++
			j++
		} else if n1[i] < n2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}
