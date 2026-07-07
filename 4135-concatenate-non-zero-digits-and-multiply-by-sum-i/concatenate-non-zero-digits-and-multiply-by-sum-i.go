func sumAndMultiply(n int) int64 {
    nonZero := 0
    sum := 0

    for n > 0 {
        mod := n%10
        if mod > 0 {
            nonZero*=10
            nonZero+=(mod)
        }

        n = n/10
        sum += mod
    }

    swap := 0
    for nonZero > 0 {
        swap*=10
        swap+=(nonZero%10)
        nonZero/=10
    }

    return int64(swap * sum)
}