package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assert := func(t testing.TB, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	}

	t.Run("Should return the sum of 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assert(t, got, want, numbers)
	})

	t.Run("Should return the sum of variable elements", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assert(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Go does not let you use equality operators with slices.
	// You could write a function to iterate over each got and want slice and check their values but for convenience sake,
	// we can use reflect.DeepEqual which is useful for seeing if any two variables are the same
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but want %v", got, want)
		}
	}

	t.Run("Sholud return an array with each sum of each slice avoiding the head of the slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("Safe mode to sum an empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 5})
		want := []int{0, 7}
		checkSums(t, got, want)
	})

}
