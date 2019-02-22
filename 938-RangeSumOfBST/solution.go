/*
 * Created on Sun Feb 10 2019
 */

package solution

import "leetcode/datastructures/tree"

func RangeSumBST(root *tree.TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if L <= root.Value && root.Value <= R {
		return root.Value + RangeSumBST(root.Left, L, R) + RangeSumBST(root.Right, L, R)
	}

	if root.Value < L {
		return RangeSumBST(root.Right, L , R)
	}

	if root.Value > R {
		return RangeSumBST(root.Left, L, R)
	}

	return 0;
}

