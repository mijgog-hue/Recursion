// LeetCode 226. Invert Binary Tree — Easy
// DFS, O(n) time, O(h) space
package main
 
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
 
type Solution struct{}
 
// InvertTree переворачивает бинарное дерево зеркально:
// для каждого узла меняет местами левое и правое поддерево.
func (s Solution) InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
 
	root.Left, root.Right = root.Right, root.Left
 
	s.InvertTree(root.Left)
	s.InvertTree(root.Right)
 
	return root
}
