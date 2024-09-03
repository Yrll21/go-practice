package iteration

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		showError(t, numbers, got, want)
	})
}

func showError(t *testing.T, numbers []int, got, want int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
