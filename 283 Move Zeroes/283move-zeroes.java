class Solution {
    public void moveZeroes(int[] nums) {
        int nextPos = 0;
        for(int i = 0; i < nums.length; i++){
            if(nums[i] != 0){
                if(nextPos < i){
                    nums[nextPos] = nums[i];
                    nums[i] = 0;
                }                
                nextPos++;
            }
        }
    }
}