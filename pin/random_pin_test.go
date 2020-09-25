package pin

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

var (
	errGetGenerated        = errors.New("GetGenerated failed")
	errGenerateOne         = errors.New("GenerateOne failed")
	errGenerateMultiple    = errors.New("GenerateMultiple failed")
	errGenerateOneThousand = errors.New("GenerateOneThousand")
	errRandom              = errors.New("Random failed")
	errIncrement           = errors.New("Increment failed")
)

func TestGeneratedPinNumber(t *testing.T) {
	p := Init()
	result := p.GenerateOne()
	if _, err := strconv.Atoi(result); err != nil {
		t.Errorf("%w, got: %s, want: number", errGenerateOne, result)
	}
}

func TestGenerateMultiple(t *testing.T) {
	p := Init()
	result := p.GenerateMultiple(10)
	if len(result) != 10 {
		t.Errorf("%w, got: %d, want: 10", errGenerateMultiple, len(result))
	}
}

func TestGetGenerated(t *testing.T) {
	p := Init()
	p.GenerateMultiple(10)
	result := p.GetGenerated()
	if len(result) != 10 {
		t.Errorf("%w, got: %d, want: 10", errGetGenerated, len(result))
	}
}

func TestGenerateThousand(t *testing.T) {
	p := Init()
	result := p.GenerateOneThousand()
	if len(result) != 1000 {
		t.Errorf("%w, got: %d, want: 1000", errGenerateOneThousand, len(result))
	}
}

func TestRandomLengthFour(t *testing.T) {
	p := Init()
	result := p.random([]int{1, 2, 3, 4}, 0)
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4}) {
		t.Errorf("%w, got: %d, want: [1 2 3 4]", errRandom, len(result))
	}
}

func TestIncrementRepeat(t *testing.T) {
	p := Init()
	result, increment := p.increment([]int{1}, 1, 1)
	if !reflect.DeepEqual(result, []int{1}) || increment != 1 {
		t.Errorf("%w, got: %d, want: [1]", errIncrement, result)
	}
}

func TestIncrementTwo(t *testing.T) {
	p := Init()
	result, increment := p.increment([]int{1}, 2, 1)
	if !reflect.DeepEqual(result, []int{1, 2}) || increment != 2 {
		t.Errorf("%w, got: %d, want: [1]", errIncrement, result)
	}
}

func TestIncrementThree(t *testing.T) {
	p := Init()
	result, increment := p.increment([]int{1}, 2, 3)
	if !reflect.DeepEqual(result, []int{1}) || increment != 3 {
		t.Errorf("%w, got: %d, want: [1]", errIncrement, result)
	}
}

func TestIncrementGeneral(t *testing.T) {
	p := Init()
	result, increment := p.increment([]int{1}, 4, 0)
	if !reflect.DeepEqual(result, []int{1, 4}) || increment != 0 {
		t.Errorf("%w, got: %d, want: [1 4]", errIncrement, result)
	}
}
