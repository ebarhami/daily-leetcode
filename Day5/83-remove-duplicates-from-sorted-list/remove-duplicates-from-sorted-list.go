/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    pointer := head 
    for pointer != nil {
        next := pointer.Next
        for next != nil && next.Val == pointer.Val {
            next = next.Next
        }
        pointer.Next = next
        pointer = next
    }

    return head
}