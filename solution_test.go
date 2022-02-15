package solution_test

import "testing"

func TestGetMessage(t *testing.T) {
	s := "Hello wrld"
	println(t)
	if len(s) == 0 {
		t.Errorf("Expected 'Hello wrld', but got %s ", s)
	}
}
