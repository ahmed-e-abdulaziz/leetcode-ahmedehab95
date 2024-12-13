func findMaxAverage(nums []int, k int) float64 {
    sum:=0
    for i:=0; i<k; i++ {
        sum+=nums[i]
    }
    maxAvg := float64(sum)/float64(k)
    for i:=1; (i+k-1)<len(nums); i++ {
        sum=sum-nums[i-1]+nums[i+k-1]
        maxAvg = math.Max(float64(sum)/float64(k), maxAvg)
    }
    return maxAvg
}