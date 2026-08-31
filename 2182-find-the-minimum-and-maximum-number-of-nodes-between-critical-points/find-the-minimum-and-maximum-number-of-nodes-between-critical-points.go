/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Point struct {
    num int
    idx int
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
    if head == nil || head.Next == nil {
        return []int{-1, -1}
    }
    critical := make([]Point, 0)

    var prev *ListNode
    curr := 0
    for head != nil {
        next := head.Next
        if prev != nil && next != nil {
            if (head.Val < prev.Val && head.Val < next.Val) || head.Val > prev.Val && head.Val > next.Val {
                critical = append(critical, Point {
                    num: head.Val,
                    idx: curr,
                })
            }
        }
        prev = head
        head = head.Next
        curr++
    }

    if len(critical) <= 1 {
        return []int{-1, -1}
    }

    minDis, maxDis := 100005, 0
    minn := minDis
    sort.Slice(critical, func(i, j int)bool{
        return critical[i].num < critical[i].num || 
            (critical[i].num == critical[i].num && critical[i].idx < critical[i].idx) 
    })

    for i, _ := range critical {
        if i > 0 {
            diff := critical[i].idx - critical[i-1].idx
            if diff < minn {
                minn = diff
                minDis = diff
            }
        }
    }
    maxDis = critical[len(critical)-1].idx - critical[0].idx
    
    return []int{minDis, maxDis}
}


/*
8, 1, 9, 10


*/