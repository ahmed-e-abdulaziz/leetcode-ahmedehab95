func longestOnes(nums []int, k int) int {
    maxRes:=0
    l, r := 0, 0
    curRes:=0
    zeroCount:=0
    for r < len(nums) {
        if nums[r] == 0 {
            zeroCount++
        }
        for zeroCount > k {
            maxRes = int(math.Max(float64(curRes), float64(maxRes)))
            if nums[l] == 0 {
                zeroCount--
            } 
            if l > r {
                r++
            }
            l++
            curRes--
        }
        r++
        curRes++
    }
    return int(math.Max(float64(curRes), float64(maxRes)))
}