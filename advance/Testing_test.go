package main

import (
	"fmt"
	"testing"
)

func Add (a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestAddTableDriven(t *testing.T) {
	test := []struct{a, b, expected int}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 0},
	}
	for _, tt := range test {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, result)
		}
	}
}

func TestAddSubtests(t *testing.T) {
	tests := []struct {a, b, expected int}{
		{2, 3, 5},
		{1, 1, 2},
		{0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d", tt.expected, result)
			}
		})
	}
}
