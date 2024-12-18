var letterMap map[byte][]byte
var d string
func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    d=digits
    letterMap = map[byte][]byte{
        '2': []byte{'a', 'b', 'c'},
        '3': []byte{'d', 'e', 'f'},
        '4': []byte{'g', 'h', 'i'},
        '5': []byte{'j', 'k', 'l'},
        '6': []byte{'m', 'n', 'o'},
        '7': []byte{'p', 'q', 'r', 's'},
        '8': []byte{'t', 'u', 'v'},
        '9': []byte{'w', 'x', 'y', 'z'},
    }
    ans := []string{}
    comb([]byte{}, 0, &ans)
    return ans
}

func comb(path []byte, idx int, ans *[]string) {
    if len(path)==len(d) {
        *ans = append(*ans, string(path[:]))
        return
    }

    for i:=idx; i<len(d); i++ {
        for _,l := range letterMap[d[i]] {
            path = append(path, l)
            comb(path, i+1, ans)
            path = path[:len(path)-1]
        }
    }
}