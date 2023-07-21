func maxProfit(prices []int) int {
    maks := 0
    minn := prices[0]

    for i:=1;i<len(prices);i++{
        if prices[i] - minn > maks {
            maks = prices[i] - minn
        }
        if prices[i] < minn {minn = prices[i]}
    }

    return maks
}