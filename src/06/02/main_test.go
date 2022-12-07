package main

import "testing"

func TestFindMarker1(t *testing.T) {
	// Arrange
	expected := 19
	// Act
	actual := findMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker2(t *testing.T) {
	// Arrange
	expected := 23
	// Act
	actual := findMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker3(t *testing.T) {
	// Arrange
	expected := 23
	// Act
	actual := findMarker("nppdvjthqldpwncqszvftbrmjlhg")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker4(t *testing.T) {
	// Arrange
	expected := 29
	// Act
	actual := findMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker5(t *testing.T) {
	// Arrange
	expected := 26
	// Act
	actual := findMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
