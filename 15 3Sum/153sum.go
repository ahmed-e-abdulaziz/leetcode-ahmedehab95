func threeSum(nums []int) [][]int {
    result := [][]int{}

    sort.Ints(nums)

    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] >= 1 {
            break
        }
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        for l, r := i + 1, len(nums) - 1; l < r; {
            if cur := nums[i] + nums[l] + nums[r]; cur == 0 {
                result = append(result, []int{nums[i], nums[l], nums[r]})
                for l++; l < r && nums[l] == nums[l - 1]; l++ {}
                for r--; l < r && nums[r] == nums[r + 1]; r-- {}
            } else if cur > 0 {
                r--
            } else {
                l++
            }
        }
    }


    return result
}