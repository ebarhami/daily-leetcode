/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }

    pointer := head

    for pointer != nil {
        next := pointer.Next
        for next != nil && next.Val == val {
            next = next.Next
        }
        pointer.Next = next
        pointer = pointer.Next
    }

    return head
}