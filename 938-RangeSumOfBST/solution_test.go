/*
 * Created on Sun Feb 10 2019
 *
 */
package solution

import (
	"leetcode/datastructures/tree"
	"testing"
)

func TestRangeSumBST(t *testing.T) {

	on := &tree.TreeNode{1, nil, nil}
	fr := &tree.TreeNode{4, nil, nil}
	sn := &tree.TreeNode{7, nil, nil}
	sx := &tree.TreeNode{6, fr, sn}
	th := &tree.TreeNode{3, on, sx}
	thr := &tree.TreeNode{13, nil, nil}
	frt := &tree.TreeNode{14, thr, nil}
	ten := &tree.TreeNode{10, nil, frt}
	et := &tree.TreeNode{8, th, ten}


	var tests =[]struct {
		a,b int
		sum int
	}{
		{1, 7, 21},
		{7, 13, 38},
		{13, 14, 27},
		{4, 6, 10},
		{8, 13, 31},
		{8, 8, 8},
	}

	for _, test := range tests {
		if res := RangeSumBST(et, test.a, test.b); res!=test.sum {
			t.Errorf("RangeSumBST(%d, %d) = %v", test.a, test.b, res)
		}
	}
}