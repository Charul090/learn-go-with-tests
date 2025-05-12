package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		get := Sum(numbers)
		want := 15
		if get != want {
			t.Errorf("expected '%d', got '%d' given %v", want, get, numbers)
		}
	})

	t.Run("sum of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		get := Sum(numbers)
		want := 6
		if get != want {
			t.Errorf("expected '%d', got '%d' given %v", want, get, numbers)
		}
	})
}
