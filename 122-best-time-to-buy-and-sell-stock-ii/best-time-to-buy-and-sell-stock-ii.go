func maxProfit(prices []int) int {
    curr := prices[0]
    profit := 0
    for i, price := range prices {
        if i == 0 {continue}
        if price < curr {
            curr = price
        } else if price > curr {
            profit += (price-curr)
            curr = price
        }
    }
    return profit
}

/*
[7,1,5,3,6,4]
curr : 7
price = 1

curr = 1
price 5

*/