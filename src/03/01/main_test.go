package main

import "testing"

func TestProcess(t *testing.T) {
	// Arrange
	expected := 157
	// Act
	actual := process("test.txt")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
