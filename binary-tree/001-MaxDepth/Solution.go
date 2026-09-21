№104 — Maximum Depth of Binary Tree
Time Complexity: O(n) Space Complexity: O(h)
package main

// type TreeNode struct { 
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
// }        


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
