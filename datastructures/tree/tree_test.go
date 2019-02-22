/**
 * File              : tree_test.go
 * Author            : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 * Date              : 13.02.2019
 * Last Modified Date: 13.02.2019
 * Last Modified By  : Ravi Jadhav <ravi.rjadhav9@gmail.com>
 */

package tree

import (
	"reflect"
	"testing"
)

func TestTreeInorder(t *testing.T) {

	on := &TreeNode{1, nil, nil}
	fr := &TreeNode{4, nil, nil}
	sn := &TreeNode{7, nil, nil}
	sx := &TreeNode{6, fr, sn}
	th := &TreeNode{3, on, sx}
	thr := &TreeNode{13, nil, nil}
	frt := &TreeNode{14, thr, nil}
	ten := &TreeNode{10, nil, frt}
	et := &TreeNode{8, th, ten}


	var tests =[]struct {
		root *TreeNode
		inorder []int
	}{
		{et,[]int{1,3,4,6,7,8,10,13,14}},
	}

	for _, test := range tests {
		if res := test.root.Inorder(); !reflect.DeepEqual(res, test.inorder) {
			t.Errorf("inorder(%q) = %v", test.inorder, res)
		}
	}
}
func TestTreeInsertion(t *testing.T) {

	et := &TreeNode{8, nil, nil}
	bTree := &BSTree{Root:et}


	var tests =[]struct {
		root *BSTree
		insert int
		inorder []int
	}{
		{bTree,3, []int{3,8}},
		{bTree,10, []int{3,8,10}},
		{bTree,1, []int{1,3,8,10}},
		{bTree,6, []int{1,3,6,8,10}},
		{bTree,14, []int{1,3,6,8,10,14}},
		{bTree,4, []int{1,3,4,6,8,10,14}},
		{bTree,7, []int{1,3,4,6,7,8,10,14}},
		{bTree,13, []int{1,3,4,6,7,8,10,13,14}},
	}

	for _, test := range tests {
		test.root.Insert(test.insert)
		if res := test.root.Inorder(); !reflect.DeepEqual(res, test.inorder) {
			t.Errorf("insertion(%q) = %v", test.inorder, res)
		}
	}
}

func TestTreeLCA(t *testing.T) {

	on := &TreeNode{1, nil, nil}
	fr := &TreeNode{4, nil, nil}
	sn := &TreeNode{7, nil, nil}
	sx := &TreeNode{6, fr, sn}
	th := &TreeNode{3, on, sx}
	thr := &TreeNode{13, nil, nil}
	frt := &TreeNode{14, thr, nil}
	ten := &TreeNode{10, nil, frt}
	et := &TreeNode{8, th, ten}


	var tests =[]struct {
		a, b int
		root *TreeNode
		expected *TreeNode
	}{
		{1, 14, et, et},
		{1, 7, et, th},
		{13, 14, et, frt},
		{7, 4, et, sx},
		{7, 13, et, et},

	}

	for _, test := range tests {
		if res := test.root.Lca(test.a, test.b); res!=test.expected {
			t.Errorf("lca(%q) = %v", test.expected, res)
		}
	}
}

func TestTreePath(t *testing.T) {

	on := &TreeNode{1, nil, nil}
	fr := &TreeNode{4, nil, nil}
	sn := &TreeNode{7, nil, nil}
	sx := &TreeNode{6, fr, sn}
	th := &TreeNode{3, on, sx}
	thr := &TreeNode{13, nil, nil}
	frt := &TreeNode{14, thr, nil}
	ten := &TreeNode{10, nil, frt}
	et := &TreeNode{8, th, ten}


	var tests =[]struct {
		a int
		root *TreeNode
		sum int
	}{
		{1, et, 12},
		{7, et, 24},
		{14, et, 32},
		{4, et, 21},
		{13, et, 45},

	}

	for _, test := range tests {
		if res := test.root.FindPathSum(test.a); res!=test.sum {
			t.Errorf("FindPathSum(%q) = %v", test.sum, res)
		}
	}
}