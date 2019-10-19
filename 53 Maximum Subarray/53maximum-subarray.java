class Solution {
    public int maxSubArray(int[] nums) {
        int curMaxSum = nums[0];
        int maxSum = nums[0];
        for(int i=1; i<nums.length; i++){
            curMaxSum = Math.max(nums[i], curMaxSum+nums[i]);
            if(maxSum < curMaxSum) maxSum = curMaxSum;
        }
        return maxSum;
    }
}