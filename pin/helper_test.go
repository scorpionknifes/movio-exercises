package pin

import (
	"errors"
	"testing"
)

var errSliceToString = errors.New("sliceToString was incorrect")

func TestSliceToStringGeneral(t *testing.T) {
	result := sliceToString([]int{1, 2, 3, 4})
	if result != "1234" {
		t.Errorf("%w, got: %s, want: 1234", errSliceToString, result)
	}
}

func TestSliceToStringZeroInfront(t *testing.T) {
	result := sliceToString([]int{0, 2, 3, 4})
	if result != "0234" {
		t.Errorf("%w, got: %s, want: 0234", errSliceToString, result)
	}
}
