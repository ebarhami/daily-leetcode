func peakIndexInMountainArray(arr []int) int {
    l, r := 0, len(arr)-1
    maks, idx := 0, 0
    for l <= r {
        mid := (l+r)/2
        goLeft := false
        if arr[mid] > maks {
            maks = arr[mid]
            idx = mid
        }
        if mid > 0 && (arr[mid] < arr[mid-1]) {
            goLeft = true
        } else if mid < len(arr)-1 && (arr[mid] > arr[mid+1]) {
            goLeft = true
        }

        if goLeft {
            r = mid-1
        } else {
            l = mid+1
        }
    }

    return idx
}