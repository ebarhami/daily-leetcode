/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {return nil}
    if head.Next == nil {return head}
    temp := head.Next

    head.Next = swapPairs(temp.Next)
    temp.Next = head

    return temp
}

/*
1 2 3 4
2 1 3 2


1 2 3
2 1 3

*/