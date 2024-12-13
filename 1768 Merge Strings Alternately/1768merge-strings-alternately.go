import "strings"

func mergeAlternately(word1 string, word2 string) string {
    var builder strings.Builder
    var i1, i2 int
    for i1 < len(word1) || i2 < len(word2) {
        if i1 < len(word1) {
            builder.WriteString(string(word1[i1]))
            i1++
        }
        if i2 < len(word2) {
            builder.WriteString(string(word2[i2]))
            i2++
        }
    }
    return builder.String()
}