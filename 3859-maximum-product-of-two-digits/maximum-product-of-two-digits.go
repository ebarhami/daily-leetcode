func maxProduct(n int) int {
    first, second := 0, 0
    for n > 0 {
        val := n % 10
        n /= 10

        if val > first {
            if first > second {
                second = first
            }
            first = val

        } else if val > second {
            second = val
        }
    }

    return first * second
}