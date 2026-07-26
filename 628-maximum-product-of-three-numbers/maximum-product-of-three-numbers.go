func maximumProduct(nums []int) int {
    first, second, third := -1001, -1001, -1001
    minFirst, minSec := 1001, 1001

    for _, num := range nums {
        if num < minFirst {
            minSec = minFirst
            minFirst = num
        } else if num < minSec {
            minSec = num
        }

        if num > first {
            third = second
            second = first
            first = num 
        } else if num > second {
            third = second
            second = num
        } else if num > third {
            third = num
        }
    }

    pos := first * second * third
    neg := first * minFirst * minSec
    if pos > neg {
        return pos
    }
    return neg
}