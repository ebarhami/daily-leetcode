/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var l, r, lPointer, rPointer *ListNode

    for head != nil {
        if head.Val < x {
            if l == nil {
                l = head
                lPointer = l
            } else {
                lPointer.Next = head
                lPointer = head
            }
        } else {
            if r == nil {
                r = head
                rPointer = r
            } else {
                rPointer.Next = head
                rPointer = head
            }
        }

        head = head.Next
    }
    if rPointer != nil {
        rPointer.Next = nil
    }
    if lPointer != nil {
        lPointer.Next = r
        return l
    } else {
        return r
    }
}