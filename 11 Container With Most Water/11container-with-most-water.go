func init() {
    debug.SetGCPercent(0)
}

func maxArea(height []int) int {
    max, l, h := 0, 0, len(height) - 1
    for l < h {
        a := (h-l) * Min(height[h], height[l])
        max = Max(max, a)
        if height[l] < height[h] {
            l++
        } else {
            h--
        }
    }
    return max
}

func Max(a, b int) int{
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int{
	if a < b {
		return a
	} else {
		return b
	}
}