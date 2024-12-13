import "strings"

func longestCommonPrefix(strs []string) string {
    var builder strings.Builder
    minLen := len(strs[0])
    for i:=0; i<minLen; i++ {
        isPrefix := true
        for j:=1; j<len(strs) && isPrefix; j++ {
            if len(strs[j]) == 0 {
                isPrefix = false
                break
            }
            if strs[j][i] != strs[j-1][i] {
                isPrefix = false
            }
            if len(strs[j]) < minLen {
                minLen = len(strs[j])
            }
        }
        if !isPrefix {
            i = minLen
        } else {
            builder.WriteString(string(strs[0][i]))
        }
    }
    return builder.String()
}