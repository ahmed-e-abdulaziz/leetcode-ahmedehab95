func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	needCount := len(tCount)
	haveCount := 0
	left := 0
	idxResultWindow := [2]int{0, math.MaxInt64}
	windowCount := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		char := s[right]
		windowCount[char]++

		if windowCount[char] == tCount[char] {
			haveCount++
		}

		for haveCount == needCount {
			if right-left < idxResultWindow[1]-idxResultWindow[0] {
				idxResultWindow = [2]int{left, right}
			}

			windowCount[s[left]]--
			if windowCount[s[left]] < tCount[s[left]] {
				haveCount--
			}
			left++
		}
	}

	if idxResultWindow[1] == math.MaxInt64 {
		return ""
	}
	return s[idxResultWindow[0] : idxResultWindow[1]+1]
}