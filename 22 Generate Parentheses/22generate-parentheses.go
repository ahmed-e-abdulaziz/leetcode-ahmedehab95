func generateParenthesis(n int) []string {
    ans := []string{}
    generate(n, 1, 0, "(", &ans)
    return ans
}

func generate(max, leftCount, rightCount int, comb string, ans *[]string) {
    if len(comb) == (max*2) {
        *ans = append(*ans, comb)
        return
    }
    if leftCount < max {
        generate(max, leftCount+1, rightCount, comb+"(", ans)
    }
    if rightCount < leftCount {
        generate(max, leftCount, rightCount+1, comb+")", ans)
    }
}