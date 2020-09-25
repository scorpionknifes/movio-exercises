package merge

import (
	"reflect"
	"testing"
)

func TestSimpleMerge(t *testing.T) {
	a := []int{1, 1, 3, 5}
	b := []int{1, 2, 3, 4}

	expected := []int{1, 1, 1, 2, 3, 3, 4, 5}
	actual := SortedMerge(a, b)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Merge was incorrect, got: %v, want: %v.", actual, expected)
	}
}
