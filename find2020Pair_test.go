package integers

import "testing"

func TestFind2020Pair(t *testing.T) {
	result := Find2020Pair(2000, 20)
	expected := true

	if result != expected {
		t.Errorf("expected '%v' but got '%v'", expected, result)
	}
}
