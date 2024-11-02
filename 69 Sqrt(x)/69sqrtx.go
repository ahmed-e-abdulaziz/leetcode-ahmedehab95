func mySqrt(x int) int {
    start, end := 0, x+1
    for start < end {
        mid := start + (end-start)/2
        if mid*mid > x {
            end = mid
        } else {
            start = mid + 1
        }
    }
    return start - 1
}
