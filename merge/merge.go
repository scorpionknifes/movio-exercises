package merge

// SortedMerge sort two int slice into one
func SortedMerge(a []int, b []int) []int {
	sorted := make([]int, len(a)+len(b))
	posA, posB, i := 0, 0, 0

	for len(a) > posA && len(b) > posB {
		if a[posA] < b[posB] {
			sorted[i] = a[posA]
			posA++
		} else {
			sorted[i] = b[posB]
			posB++
		}
		i++
	}

	for j := posA; j < len(a); j++ {
		sorted[i] = a[j]
		i++
	}
	for j := posB; j < len(b); j++ {
		sorted[i] = b[j]
		i++
	}

	return sorted
}
