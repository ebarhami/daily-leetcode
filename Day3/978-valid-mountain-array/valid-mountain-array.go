func validMountainArray(arr []int) bool {
    n := len(arr)
    if n < 3 {return false}

    isUp := true

    for i:=1;i<n;i++{
        if arr[i] > arr[i-1] { // up
            if !isUp {
                return false
            }
        } else if arr[i] < arr[i-1] { // down
            if i == 1 {
                return false
            }
            isUp = false
        } else { // same
            return false
        }
    }

    return !isUp
}
