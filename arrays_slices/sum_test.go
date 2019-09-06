package arrays_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArray(numbers)
	want := 15

	assert.Equal(t, got, want)
}

func TestSliceSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlice(numbers)
		want := 15

		assert.Equal(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		assert.Equal(t, got, want)

		numbers = []int{3, 4}

		got = SumSlice(numbers)
		want = 7

		assert.Equal(t, got, want)
	})
}
