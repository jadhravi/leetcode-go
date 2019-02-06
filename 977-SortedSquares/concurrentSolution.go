/**
 * File              : concurrentSolution.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 05.02.2019
 * Last Modified Date: 05.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */

package solution

func ConcurrentSortedSquares(A []int) []int {
	res := make([]int, len(A))
	ch := make(chan int)
	go squareIt(A, ch)
	i := len(A) - 1
	for val := range ch {
		res[i] = val
		i--
	}
	return res
}

func squareIt(A []int, ch chan int) {
	left, right := 0, len(A)-1
	var lsq, rsq int
	for left != right {
		lsq = A[left] * A[left]
		rsq = A[right] * A[right]
		if lsq <= rsq {
			ch <- rsq
			right--
		} else {
			ch <- lsq
			left++
		}
	}
	ch <- A[left] * A[left]
	close(ch)
}
