package main

import "testing"

func TestProcess(t *testing.T) {
	// Arrange
	expected := 2
	// Act
	actual := process("test.txt")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func FinalProcess(t *testing.T) {
	// Arrange
	expected := 490
	// Act
	actual := process("final.txt")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
