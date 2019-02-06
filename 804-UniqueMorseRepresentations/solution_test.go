/**
 * File              : solution_test.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */

package solution

import "testing"

func TestMorseRepresentation(t *testing.T) {
	var tests = []struct {
		input  []string
		output int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
	}

	for _, test := range tests {
		if res := UniqueMorseCodeRepresentations(test.input); res != test.output {
			t.Errorf("UniqueMorseCodeRepresentation(%q) = %v", test.input, res)
		}
	}
}
