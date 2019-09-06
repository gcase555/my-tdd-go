package arrays_slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraySum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := SumArray(numbers)
	want := 15

	assert.Equal(t, want, got)
}

func TestSliceSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlice(numbers)
		want := 15

		assert.Equal(t, want, got)
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		assert.Equal(t, got, want)

		numbers = []int{3, 4}

		got = SumSlice(numbers)
		want = 7

		assert.Equal(t, want, got)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	assert.Equal(t, want, got)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assert.Equal(t, want, got)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assert.Equal(t, want, got)
	})
}
