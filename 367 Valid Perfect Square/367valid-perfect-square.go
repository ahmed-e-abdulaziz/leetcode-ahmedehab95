func isPerfectSquare(num int) bool {
    if num == 1 {
        return true
    }
    l, r := 0, num
    for r >= l {
        mid := (l + r) / 2
        if mid*mid > num {
            r = mid - 1
        } else if mid*mid < num {
            l = mid + 1
        } else {
            return true
        }
    }
    return false
}