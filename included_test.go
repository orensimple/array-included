package included_test

import (
	included "array-included"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInclude(t *testing.T) {
	var emptyArray []int

	testArray := []int{1, 2, 3, 3, 5, 7, 9, 11, 12, 99}
	tests := []struct {
		name     string
		array    []int
		arrayInc []int
		expected bool
	}{
		{name: "Simple with repeat", array: testArray, arrayInc: []int{2, 3, 3}, expected: true},
		{name: "Simple with repeat", array: testArray, arrayInc: []int{3, 3, 5}, expected: true},
		{name: "Simple", array: testArray, arrayInc: []int{7, 9, 11}, expected: true},
		{name: "Simple", array: testArray, arrayInc: []int{2, 3, 11}, expected: false},
		{name: "Empty array", array: emptyArray, arrayInc: []int{3, 5, 7}, expected: false},
		{name: "Empty array included", array: testArray, arrayInc: emptyArray, expected: true},
		{name: "Simple one element array included", array: testArray, arrayInc: []int{11}, expected: true},
		{name: "Simple one element array included", array: testArray, arrayInc: []int{13}, expected: false},
		{name: "With long array included", array: testArray, arrayInc: []int{1, 2, 3, 3, 5, 7, 9, 11}, expected: true},
		{name: "With long array included", array: testArray, arrayInc: []int{1, 2, 3, 3, 5, 7, 9, 10}, expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := included.IsInclude(tc.array, tc.arrayInc)
			require.Equal(t, tc.expected, result)
		})
	}
}
