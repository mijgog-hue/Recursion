package main

// type TreeNode struct { 
//    Val   int
//    Left  *TreeNode
//    Right *TreeNode
// }        


func (s Solution) Count_leaves(root *TreeNode) int {
    
    if root ==nil{
        return 0
    } 

if root.Left == nil && root.Right == nil{
    return 1
}
   left := s.Count_leaves(root.Left)
    right := s.Count_leaves(root.Right)

	return right + left
}






Barglar soni
Berilgan daraxtning nechta bargi borligini toping.

Ketma-ket turgan elementlarni tepadan pastga qaraganda, ota va bola deydigan bo'lsak, bolasi yo'q elementlar barglar deyiladi.

Misol 1:
binary tree 1

Kiritma: [1, 2, 3]
Natija: 2
Misol 2:
binary tree 2

Kiritma: [5, 1, 4, null, null, 3, 6]
Natija: 3
