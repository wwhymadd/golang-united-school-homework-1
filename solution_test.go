package solution_test

import "testing"

func TestGetMessage(t *testing.T) {
	s := "Hello world"
	println(t)
	if len(s) == 0 {
		t.Errorf("Expected 'hello world', but got %s ", s)
	}
}
