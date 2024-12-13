func maxProfit(prices []int) int {
    min := math.MaxInt
    res := 0
    for i:=1; i<len(prices); i++ {
        if min > prices[i-1] {
            min = prices[i-1]
        }
        possibleRes := prices[i] - min
        if possibleRes > res {
            res = possibleRes
        }
    }
    return res
}