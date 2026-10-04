package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("fixed size array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		checkInt(t, got, want, numbers[:])
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum2(numbers)
		want := 6

		checkInt(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	checkSlices(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSlices(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSlices(t, got, want)
	})
}

func checkSlices(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkInt(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d but got %d, given %v", want, got, numbers)
	}
}
