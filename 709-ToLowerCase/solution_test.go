/**
 * File              : solution_test.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */
package solution

import "testing"

func TestToUpperCase(t *testing.T) {
	var tests = []struct {
		input, output string
	}{
		{"TESTING", "testing"},
		{"testing", "testing"},
		{"TestIng", "testing"},
		{"Test[ng#2", "test[ng#2"},
	}

	for _, test := range tests {
		if res := ToLowerCase(test.input); res != test.output {
			t.Errorf("ToLowerCase(%q) = %v", test.input, res)
		}
	}
}
