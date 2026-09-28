LeetCode 543 — Diameter of Binary Tree
Time: O(n²)
Space: O(h)
package main

// type TreeNode struct { 
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
// }        
func (s Solution) max(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := s.max(root.Left)
    right := s.max(root.Right)

    if left > right {
        return left + 1
    }

    return right + 1
}

func (s Solution) DiameterOfBinaryTree(root *TreeNode) int {
    if root == nil{
        return 0
    }
    
    result := 0

left := s.max(root.Left)
right := s.max(root.Right)

    current := left + right

if current > result {
    result = current
}
LeftHigh := s.DiameterOfBinaryTree(root.Left)
RightHigh := s.DiameterOfBinaryTree(root.Right)

if LeftHigh > result{
    result = LeftHigh
}
if RightHigh > result {
    result = RightHigh
}
return result
}
