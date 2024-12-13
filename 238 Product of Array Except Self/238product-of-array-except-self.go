func productExceptSelf(nums []int) []int {
    length := len(nums)
    left, right := make([]int, length), make([]int, length)
    left[0] = 1
    for i := 1; i < length; i++ {
        left[i] = nums[i-1] * left[i-1]
    }

    right[length-1]=1
    for i := length-2; i>=0; i-- {
        right[i] = nums[i+1] * right[i+1] 
    }
    
    ans := make([]int, length)
    for i:=0; i<len(nums); i++ {
        ans[i] = left[i] * right[i]
    }
    return ans
}