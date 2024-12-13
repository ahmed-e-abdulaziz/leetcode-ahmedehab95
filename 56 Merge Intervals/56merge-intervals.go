func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    ans := [][]int{intervals[0]}
    for i:=1; i<len(intervals); i++ {
        ansStart, ansEnd, curStart, curEnd := ans[len(ans)-1][0], ans[len(ans)-1][1], intervals[i][0], intervals[i][1]
        if ansEnd >= curStart {
            if ansStart > curStart {
                ans[len(ans)-1][0] = curStart
            }
            if ansEnd < curEnd {
                ans[len(ans)-1][1] = curEnd
            }
        } else {
            ans = append(ans, intervals[i])
        }
    }
    return ans
}
