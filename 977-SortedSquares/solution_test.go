/**
 * File              : solution_test.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */
package solution

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		input, output []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, test := range tests {
		if res := SortedSquares(test.input); !reflect.DeepEqual(res, test.output) {
			t.Errorf("SortedSquares(%v) = %v", test.input, res)
		}
	}
}

func TestConcurrentSortedSquares(t *testing.T) {
	var tests = []struct {
		input, output []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, test := range tests {
		if res := ConcurrentSortedSquares(test.input); !reflect.DeepEqual(res, test.output) {
			t.Errorf("ConcurrentSortedSquares(%v) = %v", test.input, res)
		}
	}
}
