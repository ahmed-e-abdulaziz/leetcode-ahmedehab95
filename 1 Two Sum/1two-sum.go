func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		numsMap[num] = i
	}
	for i, num := range nums {
		diff := target - num
		if numsMap[diff] == i {
			continue
		}
		j, ok := numsMap[diff]
		if ok {
			return []int{i, j}
		}
	}
	return []int{-1, -1}
}
