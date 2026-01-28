func maxProfit(prices []int) int {
    ans := 0
    cur := prices[0]
    for _, price := range prices[1:] {
        if cur < price {
            ans = max(ans, price - cur)
        } else {
            cur = price
        }
    }
    return ans
}