/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    fast, slow := &ListNode{Next: head}, &ListNode{Next: head}
    answer := slow

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = slow.Next.Next

    return answer.Next
}