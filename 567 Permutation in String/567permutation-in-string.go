func checkInclusion(s1 string, s2 string) bool {
    count, need:=prepF(s1)
    window:=make([]int, 123, 123)
    have:=0
    l,r:=0,0
    for r < len(s2) {
        if count[s2[r]] > 0 && window[s2[r]] < count[s2[r]] {
            window[s2[r]]++
            if window[s2[r]] == count[s2[r]] {
                have++
            }
            if have == need {
                return true
            }
            r++
        } else {
            window[s2[l]]--
            if count[s2[l]] > 0 && window[s2[l]] == count[s2[l]]-1 {
                have--
            }
            l++
            if l > r {
                r++
            }
        }
    }
    return false
}

func prepF(s string) ([]int, int) { // Count letters' frequency in the target string
	count := make([]int, 123, 123)
    set := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
        set[s[i]]=true
	}
	return count, len(set)
}
