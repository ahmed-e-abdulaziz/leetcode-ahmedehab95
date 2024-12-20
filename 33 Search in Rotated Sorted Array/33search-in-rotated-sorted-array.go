func search(nums []int, target int) int {
    min:=findMin(nums)
    var l, r int
    
    if target >= nums[0] {
        l, r = 0, min - 1
        if r == - 1{
            r = len(nums) - 1
        }
    } else {
        l, r = min, len(nums)-1
    }

    for l <= r {
        mid := (l+r)/2
        if nums[mid] < target {
            l = mid + 1
        } else if nums[mid] > target{
            r = mid - 1
        } else {
            return mid
        }
    }
    return -1
}

func findMin(nums []int) int {
    l, r:=0, len(nums)-1
    mid:=0
    for l < r {
        mid = (l+r)/2
        if nums[mid] > nums[r] {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l
}