package packs_test

import (
	"gordon-to/re-golang-demo/internal/packs"
	"testing"
	"reflect"
)

func TestItemsToPacks(t *testing.T) {
	testCases := []struct {
		items          int
		packs          []int
		expectedResult map[int]int
		expectedOverOrder int
		expectedError  error
	}{
		{
			1,
			[]int{500},
			map[int]int{500: 1},
			499,
			nil,
		},
		{
			5000,
			[]int{500, 2000, 1000},
			map[int]int{500: 0, 2000: 2, 1000: 1},
			0,
			nil,
		},
	}
	for _, testCase := range testCases {
		result, overOrder, err := packs.ItemsToPacks(testCase.items, testCase.packs)
		if err != testCase.expectedError {
			t.Errorf("Expected %v, but got %v", testCase.expectedError, err)
		}
		if overOrder != testCase.expectedOverOrder {
			t.Errorf("Expected %v, but got %v", testCase.expectedOverOrder, overOrder)
		}
		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Errorf("Expected %v, but got %v", testCase.expectedResult, result)
		}
	}
}
