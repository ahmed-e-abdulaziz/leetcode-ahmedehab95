func romanToInt(s string) int {
    res := 0
    romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    for i := 0; i < len(s); i++ {
        res+=romanMap[s[i]]
        if i < (len(s) - 1) && s[i] == 'C' && (s[i + 1] == 'M' || s[i + 1] == 'D') {
            if s[i + 1] == 'M'{
                res+=800
            }
            if s[i + 1] == 'D' {
                res+=300
            }
            i++
        }
        if i < len(s) - 1 && s[i] == 'X' && (s[i + 1] == 'L' || s[i + 1] == 'C') {
            if s[i + 1] == 'C'{
                res+=80
            }
            if s[i + 1] == 'L' {
                res+=30
            }
            i++
        }
        if i < (len(s) - 1) && s[i] == 'I' && (s[i + 1] == 'V' || s[i + 1] == 'X') {
            if s[i + 1] == 'X'{
                res+=8
            }
            if s[i + 1] == 'V' {
                res+=3
            }
            i++
        }
    }
    return res
}