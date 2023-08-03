/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {return nil}
    
    for len(lists) > 1 {
        newLists := make([]*ListNode, 0)
        for i:=0;i<len(lists);i+=2{
            first := lists[i]
            var second *ListNode
            if i+1 < len(lists) {
                second = lists[i+1]
            }
            newLists = append(newLists, mergeTwoLists(first,second))
        }
        lists = newLists
    }
    

    return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {return list2}
    if list2 == nil {return list1}

    if list1.Val < list2.Val {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    } else {
        list2.Next = mergeTwoLists(list1, list2.Next)
        return list2
    }
}