package pin

import "strconv"

// sliceToString - convert a slice int to a string pin
func sliceToString(a []int) string {
	b := ""
	for _, v := range a {
		b += strconv.Itoa(v)
	}
	return b
}
