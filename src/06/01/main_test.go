package main

import "testing"

func TestFindMarker1(t *testing.T) {
	// Arrange
	expected := 7
	// Act
	actual := findMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker2(t *testing.T) {
	// Arrange
	expected := 5
	// Act
	actual := findMarker("bvwbjplbgvbhsrlpgdmjqwftvncz")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker3(t *testing.T) {
	// Arrange
	expected := 6
	// Act
	actual := findMarker("nppdvjthqldpwncqszvftbrmjlhg")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker4(t *testing.T) {
	// Arrange
	expected := 10
	// Act
	actual := findMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestFindMarker5(t *testing.T) {
	// Arrange
	expected := 11
	// Act
	actual := findMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	// Assert
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}


