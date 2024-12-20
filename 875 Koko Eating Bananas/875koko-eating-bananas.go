func minEatingSpeed(piles []int, h int) int {
	l, r := 1, slices.Max(piles)
	for l < r {
		mid := (l + r) / 2
		if isValid(piles, h, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isValid(piles []int, maxHours int, bananaEatingSpeed int) bool {
	for _, bananaPile := range piles {
		hoursNeededToEatThePile := ((bananaPile - 1) / bananaEatingSpeed) + 1
		maxHours -= hoursNeededToEatThePile
		if maxHours < 0 {
			return false
		}
	}
	return true
}