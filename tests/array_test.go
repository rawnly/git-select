package tests

import (
	"github.com/rawnly/git-select/array"
	"testing"
)

func TestReduce(t *testing.T) {
	base := []int{1, 2, 3}

	value := array.Reduce[int, int](base, 0, func(idx int, acc int, item int) int {
		return acc + item
	})

	if value != 6 {
		t.Errorf("Expected 6 got: %d", value)
	}
}

func TestFilter(t *testing.T) {
	base := []int{2, 3}

	value := array.Filter[int](base, func(item int, i int) bool {
		return item > 2
	})

	if value[0] != 3 {
		t.Errorf("Expected 3 got: %d", value[0])
	}
}
