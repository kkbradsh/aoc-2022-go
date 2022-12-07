package main

import "testing"

func TestProcess(t *testing.T) {
	// Arrange
	expected := "MCD"
	// Act
	actual := process("test.txt")
	// Assert
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

