func lengthOfLIS(nums []int) int {
    LISArr := make([]int, 0)

    for _, num := range nums {
        if len(LISArr) == 0 || LISArr[len(LISArr)-1] < num {
            LISArr = append(LISArr, num)
        } else {
            pos := getIdx(num, LISArr)
            LISArr[pos] = num
        }
    }

    return len(LISArr)
}

func getIdx(num int, LISArr []int) int {
    l, r := 0, len(LISArr) - 1
    
    for r > l {
        mid := (l+r)/2
        if LISArr[mid] < num {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return r
}

/*

1 4 6
5 -> 1 4 5
3 -> 1 3 6

*/
