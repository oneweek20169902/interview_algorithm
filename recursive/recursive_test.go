package recursive

import (
	"math"
	"testing"
)

func TestGetMax(t *testing.T) {
	testCases := []struct {
		name        string
		input       []int
		expectedMax int
	}{
		{
			name:        "Valid positive integers",
			input:       []int{1, 3, 5, 2, 4},
			expectedMax: 5,
		},
		{
			name:        "Valid negative integers",
			input:       []int{-1, -3, -5, -2, -4},
			expectedMax: -1,
		},
		{
			name:        "Valid mix of positive and negative integers",
			input:       []int{-1, 3, 5, -2, 4},
			expectedMax: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetMax(tc.input)
			if result != tc.expectedMax {
				t.Errorf("Expected maximum value to be %d, but got %d", tc.expectedMax, result)
			}
		})
	}

	// Additional test cases for edge cases:
	// Empty slice case
	result := GetMax([]int{})
	if result != math.MinInt64 {
		t.Errorf("Expected maximum value to be %d, but got %d", math.MinInt64, result)
	}
	// Single element slice case
	result = GetMax([]int{42})
	if result != 42 {
		t.Errorf("Expected maximum value to be %d, but got %d", 42, result)
	}
}
