package array_slice

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("fixed size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(numbers)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}

func TestAllSumTails(t *testing.T) {

	checksum := func(t *testing.T, got, want []int) {
		t.Helper()

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("test case for non empty slice", func(t *testing.T) {

		got := SumTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}

		checksum(t, got, want)
	})

	t.Run("test case for empty slice", func(t *testing.T) {

		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checksum(t, got, want)
	})
}
