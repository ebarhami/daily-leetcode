func searchMatrix(matrix [][]int, target int) bool {
    l, r := 0, len(matrix)-1
    for l < r {
        mid := (l+r+1)/2
        if matrix[mid][0] == target {
            l = mid 
            break
        } else if matrix[mid][0] > target {
            r = mid - 1
        } else {
            l = mid
        }
    }
    row := l

    l, r = 0, len(matrix[row])-1
    for l < r {
        mid := (l+r)/2
        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] < target {
            l = mid + 1
        } else {
            r = mid -1
        }
    }

    return matrix[row][l] == target
}