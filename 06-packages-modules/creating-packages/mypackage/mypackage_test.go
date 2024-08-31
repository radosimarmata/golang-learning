package mypackage

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Go")
	expected := "Hello, Go"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
