func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 || len(s) == 1 {
        return len(s)
    }
    occurred := map[byte]bool{s[0]: true}
    l,r := 0,1
    max:=1
    sum:=1
    for r < len(s) {
        if occurred[s[r]] {
            delete(occurred, s[l])
            l++
            sum--
        } else {
            sum++
            occurred[s[r]] = true
            r++
        }
        max = int(math.Max(float64(sum), float64(max)))   
    }
    return max
}