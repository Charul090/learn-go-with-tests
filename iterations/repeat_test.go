package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeated("b", 5)
	want := "bbbbb"

	if repeated != want {
		t.Errorf("expected '%q', got '%q'", repeated, want)
	}
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeated("a", 5)
	}
}

func ExampleRepeat() {
	x := Repeated("x", 5)
	fmt.Println(x)
	// Output: xxxxx
}
