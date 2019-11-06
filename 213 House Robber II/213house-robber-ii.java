class Solution {
    public int rob(int[] nums) {
        if(nums.length == 0) return 0;
        if(nums.length == 1) return nums[0];
        if(nums.length == 2) return Math.max(nums[0], nums[1]);
      return Math.max(getHouseRobberMax1(Arrays.copyOfRange(nums, 0, nums.length -1)), getHouseRobberMax1(Arrays.copyOfRange(nums, 1, nums.length)));
    }
    
    private int getHouseRobberMax1(int[] nums) {
        int max = 0;
        for(int i = 0; i < nums.length; i++){
            if(i > 1) nums[i] += nums[i-2];
            if(nums[i] > max) max = nums[i];
            else nums[i] = max;
        }
        return max;
    }
}