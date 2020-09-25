package pin

import "strconv"

func sliceToString(a []int) string {
	b := ""
	for _, v := range a {
		b += strconv.Itoa(v)
	}

	return b
}
