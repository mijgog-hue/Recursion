83. Remove Duplicates from Sorted List
Time: O(n)
Space: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    left:= head

    for left!= nil && left.Next!=nil{
        if left.Val == left.Next.Val{
         left.Next = left.Next.Next
        } else{
            left = left.Next 
        }
    }
    return head
}
