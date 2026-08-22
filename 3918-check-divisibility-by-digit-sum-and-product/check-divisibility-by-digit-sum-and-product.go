func checkDivisibility(n int) bool {
    sum, product := 0, 1

    num := n
    for num > 0 {
        x := num % 10
        num /= 10

        sum += x
        product *= x
    }

    return n % (sum+product) == 0
}