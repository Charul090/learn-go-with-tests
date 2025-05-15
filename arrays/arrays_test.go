package arrays

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 5})
	want := []int{3, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected '%v', got '%v'", want, got)
	}
}

func TestSumTails(t *testing.T) {

	checkSums := func(got, want []int, t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%v', got '%v'", want, got)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}
		checkSums(got, want, t)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(got, want, t)
	})
}
