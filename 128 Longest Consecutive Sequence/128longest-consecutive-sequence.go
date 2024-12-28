func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		count := 1
		if _, ok := set[nums[i]]; ok {
			prev := nums[i] - 1
			_, hasPrev := set[prev]
			for hasPrev {
				delete(set, prev)
				prev--
				count++
				_, hasPrev = set[prev]
			}
			next := nums[i] + 1
			_, hasNext := set[next]
			for hasNext {
				delete(set, next)
				next++
				count++
				_, hasNext = set[next]
			}
		}
		maxCount = max(count, maxCount)
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
