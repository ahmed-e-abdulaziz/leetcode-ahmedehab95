var max int
var end int

func combine(n int, k int) [][]int {
    max = k
    end = n
    var ans [][]int
	backtrack(1, []int{}, &ans)
    return ans
}

func backtrack(num int, path []int, ans *[][]int) {
	if len(path) == max {
        *ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := num; i <= end; i++ {
        path = append(path, i)
        backtrack(i+1, path, ans)
        path = path[:len(path)-1]
	}
}
