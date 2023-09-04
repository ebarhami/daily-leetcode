/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {return false}
    fast, slow := head.Next, head

    for fast != nil {
        if fast == slow {
            return true
        }
        fast, slow = fast.Next, slow.Next
        if fast != nil {
            fast = fast.Next
        }
    }
    return false
}