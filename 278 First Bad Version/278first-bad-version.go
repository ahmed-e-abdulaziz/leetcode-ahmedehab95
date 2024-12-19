/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	l, r := 0, n
	for l < r-1 {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}