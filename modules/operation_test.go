package modules_test

import (
	"testing"

	"github.com/post-human-world/go-testing-for-jenkins-pipeline/modules"
)

func TestAdd(t *testing.T) {
	result := modules.Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	result := modules.Add(-2, -3)
	expected := -5

	if result != expected {
		t.Errorf("Add(-2, -3) returned %d, expected %d", result, expected)
	}
}
