/**
 * File              : solution.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */
package solution

func ToLowerCase(str string) string {
	lowerRunes := []rune(str)
	for i, value := range str {
		if isLowerCase(value) {
			lowerRunes[i] = value
			continue
		}
		lowerRunes[i] = lowerCase(value)
	}
	return string(lowerRunes)
}

func isLowerCase(r rune) bool {
	return r < 'A' || r > 'Z'
}

func lowerCase(r rune) rune {
	return r + 32
}
