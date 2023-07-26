/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    check := make(map[*ListNode]struct{})

    for headA != nil {
        check[headA] = struct{}{}
        headA = headA.Next
    }

    for headB != nil {
        if _, ok := check[headB];ok {
            return headB
        }
        headB = headB.Next
    }

    return nil
}