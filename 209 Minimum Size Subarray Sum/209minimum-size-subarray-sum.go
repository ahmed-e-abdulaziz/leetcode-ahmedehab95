func minSubArrayLen(target int, nums []int) int {
	l, r, sum, minLength := 0, 0, 0, math.MaxInt
	for r < len(nums) || (r >= len(nums) && sum >= target) {
		if sum < target {
			sum += nums[r]
			r++
		} else {
			sum -= nums[l]
			l++
			if r-l+1 < minLength {
				minLength = r - l + 1
			}
			if minLength == 1 {
				return minLength
			}
		}
	}
	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}