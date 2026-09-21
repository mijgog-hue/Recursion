LeetCode 110 — Balanced Binary Tree
Time: O(n²)
Space: O(n)
package main

/*
type TreeNode struct {
	Val int
        Left  *TreeNode
	Right *TreeNode
} */
func (s Solution) MaxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := s.MaxDepth(root.Left)
    right := s.MaxDepth(root.Right)

    if left > right {
        return left + 1
    }

    return right + 1
}



func (s Solution) HeightBalancedBinaryTree(tree *TreeNode) bool {
if tree == nil {
        return true
    }

    left := s.MaxDepth(tree.Left)
    right := s.MaxDepth(tree.Right)

if left-right > 1 || right-left >1 {
  return false
}
return s.HeightBalancedBinaryTree(tree.Left) && s.HeightBalancedBinaryTree(tree.Right)
}
