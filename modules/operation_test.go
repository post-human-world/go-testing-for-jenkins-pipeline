package modules

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	result := Add(-2, -3)
	expected := -5

	if result != expected {
		t.Errorf("Add(-2, -3) returned %d, expected %d", result, expected)
	}
}
