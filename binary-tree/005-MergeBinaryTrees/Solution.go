LeetCode 617 — Merge Two Binary Trees,
Time: O(n)
Space Complexity:
O(n)     — если считать память для результата
O(h)     — если считать только дополнительную память алгоритма

package main

/*
type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
} */

func (s Solution) MergeBinaryTrees(tree1 *TreeNode, tree2 *TreeNode) *TreeNode {
	if tree1==nil && tree2==nil{
		return nil
	}
	if tree1 == nil{
		return tree2
	}
	if tree2 == nil{
		return tree1
	}

	root := &TreeNode{
		Val: tree1.Val + tree2.Val,
	}

	root.Left = s.MergeBinaryTrees(tree1.Left,tree2.Left)
	root.Right = s.MergeBinaryTrees(tree1.Right,tree2.Right)

	return root
}
/
