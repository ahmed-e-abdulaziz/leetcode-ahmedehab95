func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap:=make([]int, 26)
    for _,c:=range s {
        sMap[c-'a']++
    }
    for _,c:=range t {
        count:=sMap[c-'a']
        if count == 0 {
            return false
        }
        sMap[c-'a']=count-1
    }
    return true
}