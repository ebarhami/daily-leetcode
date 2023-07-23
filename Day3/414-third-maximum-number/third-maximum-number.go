func thirdMax(nums []int) int {
    min := -(1<<31) -1
    a, b, c := min, min, min

    for _, num := range nums {
        if num >= a {
            if num > a {
                c = b
                b = a
            }
            a = num
        } else {
            if num >= b {
                if num > b {
                    c = b
                }
                b = num 
            } else {
                if num >= c {
                    c = num
                }
            }
        }
    }

    if c == min {
        return a
    } else {
        return c
    }
}