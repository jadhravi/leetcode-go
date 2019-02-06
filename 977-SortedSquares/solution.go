/**
 * File              : solution.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */
package solution

import "sort"

func SortedSquares(A []int) []int {
	res := make([]int, len(A))
	for i, val := range A {
		res[i] = val * val
	}
	sort.Ints(res)
	return res
}
