class Solution {
    public void rotate(int[] nums, int k) {
        int i = 0;
        int[] result = new int[nums.length];
        while(i < nums.length){
            result[(i+k)%nums.length] = nums[i];
            i++;
        }
        for(int x = 0; x < nums.length; x++)
            nums[x] = result[x];
    }
}