/**
 * File              : solution.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */
package solution

import "strings"

/**
Notes:
1. The length of words will be at most 100.
2. Each words[i] will have length in range [1, 12].
3. words[i] will only consist of lowercase letters.
**/
func UniqueMorseCodeRepresentations(words []string) int {
	if len(words) == 0 {
		return 0
	}

	rep := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	res := make(map[string]bool)
	for _, word := range words {
		var morse strings.Builder
		for _, ru := range word {
			morse.WriteString(rep[ru-'a'])
		}
		res[morse.String()] = true
	}
	return len(res)
}
