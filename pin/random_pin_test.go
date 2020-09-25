package pin

import (
	"errors"
	"strconv"
	"testing"
)

var (
	errGetGenerated        = errors.New("GetGenerated failed")
	errGenerateOne         = errors.New("GenerateOne failed")
	errGenerateMultiple    = errors.New("GenerateMultiple failed")
	errGenerateOneThousand = errors.New("GenerateOneThousand")
	errRandom              = errors.New("Random failed")
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
