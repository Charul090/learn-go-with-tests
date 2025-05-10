package integers

import "testing"

func TestAdder(t *testing.T) {
	get := Add(3, 4)
	want := 7

	if get != want {
		t.Errorf("expected '%d', got '%d'", get, want)
	}
}
