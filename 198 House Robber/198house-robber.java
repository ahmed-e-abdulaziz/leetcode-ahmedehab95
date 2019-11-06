class Solution {
    public int rob(int[] nums) {
        int max = 0;
        for(int i = 0; i < nums.length; i++){
            if(i > 1) nums[i] += nums[i-2];
            if(nums[i] > max) max = nums[i];
            else nums[i] = max;
        }
        return max;
    }
}