var t int
var nums []int
var ans [][]int
func combinationSum(candidates []int, target int) [][]int {
    t=target
    nums=candidates
    ans=[][]int{}
    comb(0, 0, []int{})
    return ans
}

func comb(index int, sum int, path []int) {
    if sum >= t {
        if sum == t {
            ans = append(ans, append([]int{}, path...))
        }
        return
    }

    for i:=index; i<len(nums); i++ {
        path = append(path, nums[i])
        comb(i, sum+nums[i], path)
        path = path[:len(path)-1]
    }
}