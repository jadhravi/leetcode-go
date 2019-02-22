/*
 * Copyright (c) 2019, Ravi Jadhav. All rights reserved.
 */

package tree

type BSTree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (bTree *BSTree ) Insert(val int) {
	bTree.Root = bTree.Root.insert(val)
}

func (bTree *BSTree ) Inorder() []int{
	return bTree.Root.Inorder()
}

func (bTree *BSTree ) Preorder() []int{
	return bTree.Root.Preorder()
}

func (bTree *BSTree ) Postorder() []int{
	return bTree.Root.Postorder()
}

func (root *TreeNode) insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: val, Left: nil, Right: nil}
	}

	if val < root.Value {
		root.Left = root.Left.insert(val)
	} else if val > root.Value {
		root.Right = root.Right.insert(val)
	}
	return root
}

func (root *TreeNode) Inorder() []int {
	if root == nil {
		return []int{}
	}

	return append(
		append(root.Left.Inorder(), root.Value),
		root.Right.Inorder()...)
}

func (root *TreeNode) Preorder() []int {
	if root == nil {
		return []int{}
	}

	return append(
		append([]int{root.Value}, root.Left.Preorder()...),
		root.Right.Preorder()...)
}

func (root *TreeNode) Postorder() []int {
	if root == nil {
		return []int{}
	}

	return append(
		append(root.Left.Postorder(), root.Right.Postorder()...),
		root.Value)
}

func (root *TreeNode) Lca(a, b int) *TreeNode {
	if root.Value == a || root.Value == b {
		return root
	}

	if (a < root.Value && b > root.Value) || (a > root.Value && b < root.Value) {
		return root
	}

	if a < root.Value && b < root.Value {
		return root.Left.Lca(a, b)
	}

	if a > root.Value && b > root.Value {
		return root.Right.Lca(a, b)
	}
	return nil
}

func (root *TreeNode) FindPathSum(val int) int {
	if root.Value == val {
		return val
	}

	if root.Value < val {
		return root.Value + root.Right.FindPathSum(val)
	}

	if root.Value > val {
		return root.Value + root.Left.FindPathSum(val)
	}

	return 0
}