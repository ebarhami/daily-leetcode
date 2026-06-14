/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    slow, fast := &ListNode{Next: head}, &ListNode{Next: head}

    for fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    // reverse slow.Next
    curr, prev, next := slow.Next, &ListNode{}, &ListNode{}
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr

        curr = next
    }

    max := -1
    for prev.Next != nil {
        sum := head.Val + prev.Val
        if sum > max {
            max = sum
        }
        head = head.Next
        prev = prev.Next
    }

    return max
}

/*
1 -> 2 -> 3




*/