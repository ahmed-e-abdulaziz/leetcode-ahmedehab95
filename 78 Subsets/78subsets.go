func subsets(nums []int) [][]int {
	return append(getSubsets(nums[0], nums[1:]), []int{})
}
func getSubsets(num int, nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{num}}
	}
	
	subsets := getSubsets(nums[0], nums[1:])
	res := [][]int{[]int{num}}
	res=append(res, subsets...)
	for _, set := range subsets {
		newGroup:=[]int{num}
		newGroup=append(newGroup, set...)
		res = append(res, newGroup)
	}
	return res
  }