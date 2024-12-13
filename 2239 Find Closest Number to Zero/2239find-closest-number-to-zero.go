import "math"

func findClosestNumber(nums []int) int {
    res := nums[0]
    for _, num := range nums {
        if math.Abs(float64(num)) < math.Abs(float64(res)) {
            res = num
        }
        if math.Abs(float64(num)) == math.Abs(float64(res)) && num > res {
            res = num
        }
    }
    return res
}