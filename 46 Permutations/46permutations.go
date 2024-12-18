func permute(nums []int) [][]int {
    ans:=[][]int{}
	backtrack(nums, []int{}, &ans)
    return ans
}

func backtrack(nums []int, path []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans=append(*ans, path)
        return
	}
	for i := 0; i < len(nums); i++ {
        // Remove your current num
		newNums := append([]int(nil), nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

        // Append your num to the new path
		newPath := append(path, nums[i])

        // Continue by backtracking
        backtrack(newNums, newPath, ans)
	}
}