package solution_test

import "testing"

func TestGetMessage(t *testing.T) {
	a := "Hello wrld"
	if len(a) == 0 {
		t.Errorf("Expected 'hello wrld', but got %s ", a)
	}
}
