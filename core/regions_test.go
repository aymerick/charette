package core

import (
	"fmt"
	"testing"
)

func TestExtractRegions(t *testing.T) {
	result := ExtractRegions("France,  World,Germany   ,Asia,Prout, USA")

	expected := []string{"France", "World", "Germany", "Asia", "USA"}

	if len(result) != len(expected) {
		t.Fatal(fmt.Sprintf("Failed to extract regions, got '%v' but expected '%v'", result, expected))
	}

	for i, value := range expected {
		if result[i] != value {
			t.Errorf("Failed to extract regions, got '%v' but expected '%v'", result, expected)
		}
	}
}
