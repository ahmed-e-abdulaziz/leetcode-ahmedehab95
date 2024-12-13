func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }
    pair := false
    start := nums[0]
    ans := []string{}
    for i := 1; i<len(nums); i++ {
        if (nums[i] - nums[i-1]) == 1 {
            pair = true
        } else {
            if pair {
                ans = append(ans, (strconv.Itoa(start) + "->" +strconv.Itoa(nums[i-1])))
            } else {
                ans = append(ans, strconv.Itoa(start))
            }
            pair = false
            start = nums[i]
        }
    }
    if pair {
        ans = append(ans, (strconv.Itoa(start) + "->" +strconv.Itoa(nums[len(nums)-1])))
    } else {
        ans = append(ans, strconv.Itoa(start))
    }
    return ans
}