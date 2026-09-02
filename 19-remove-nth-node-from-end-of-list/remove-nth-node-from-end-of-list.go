/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    sentinel := &ListNode{Next: head}
    slow, fast := sentinel, sentinel

    for range n {
        fast = fast.Next
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next

    return sentinel.Next
}

/*
n = 2
X,1,2,3,4,5]
x, 2
1, 3
2, 4
3, 5
4, nil
*/