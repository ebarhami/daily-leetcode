func peakIndexInMountainArray(arr []int) int {
    l, r, maks, idx := 0, len(arr)-1, 0, 0
    for l <= r {
        mid := (l+r)/2
        if arr[mid] > maks {
            maks = arr[mid]
            idx = mid
        }

        goLeft := (mid > 0 && (arr[mid] < arr[mid-1])) || (mid < len(arr)-1 && (arr[mid] > arr[mid+1]))
        if goLeft {
            r = mid-1
        } else {
            l = mid+1
        }
    }

    return idx
}