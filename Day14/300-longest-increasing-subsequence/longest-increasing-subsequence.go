func lengthOfLIS(nums []int) int {
    arr := make([]int, 0)

    for _, num := range nums {
        if len(arr) == 0 || arr[len(arr)-1] < num {
            arr = append(arr, num)
        } else {
            i := findPos(arr, num)
            arr[i] = num
        }
    }
    return len(arr)

}

func findPos(arr []int, x int) int {
    l, r := 0, len(arr)-1

    for l < r {
        mid := (l+r)/2
        if arr[mid] < x {
            l = mid+1
        } else if arr[mid] >= x {
            r = mid
        }
    }
    return l
}

/*
10,9,2,5,3,7,101,18, 1, 2, 3, 4, 5, 6, 7, 8

arr [1, 2, 3, 18,]



*/

