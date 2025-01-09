func prefixCount(words []string, pref string) int {
    count:=0
    for _, w := range words {
        i:=0
        valid:=true
        for i<len(pref) && i<len(w) && valid {
            if pref[i] != w[i] {
                valid = false
                continue
            }
            i++
        }
        if i == len(pref) {
            count++
        }
    }
    return count
}