LeetCode 101 — Symmetric Tree
Time: O(n)
Space: O(h), worst case O(n)

package main

// type TreeNode struct { 
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
// }        

func Check(left,right *TreeNode) bool {

if left ==nil && right ==nil{
  return true
}
if left == nil || right == nil{
  return false
}
if left.Val != right.Val{
  return false
}
return Check(left.Left,right.Right) && Check(left.Right,right.Left)
}

func (s Solution) IsSymmetric(root *TreeNode) bool {
    if root == nil{
      return true
    }
    return Check(root.Left,root.Right )

}
