func nodesBetweenCriticalPoints(head *ListNode) []int {
    if head == nil || head.Next == nil {
        return []int{-1, -1}
    }

    critical := make([]int, 0)

    var prev *ListNode
    curr := 0

    for head != nil {
        next := head.Next

        if prev != nil && next != nil {
            if (head.Val < prev.Val && head.Val < next.Val) ||
                (head.Val > prev.Val && head.Val > next.Val) {
                critical = append(critical, curr)
            }
        }

        prev = head
        head = head.Next
        curr++
    }

    if len(critical) <= 1 {
        return []int{-1, -1}
    }

    minDis := 100005

    for i := 1; i < len(critical); i++ {
        diff := critical[i] - critical[i-1]
        if diff < minDis {
            minDis = diff
        }
    }

    maxDis := critical[len(critical)-1] - critical[0]

    return []int{minDis, maxDis}
}