/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    traverse := head
    mid := middle(head)

    rev := reverse(mid)

    for rev != nil {
        if rev.Val != traverse.Val{
            return false
        }
        rev, traverse = rev.Next, traverse.Next
    }

    return true
}

func middle(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast != nil {
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
        slow = slow.Next
    }

    return slow
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        temp := head.Next
        head.Next = prev
        prev = head
        head = temp
    }

    return prev
}
