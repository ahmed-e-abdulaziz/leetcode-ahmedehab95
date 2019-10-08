class Solution {
    public int removeDuplicates(int[] nums) {
        int size = 0;
        int current = -2147483648;
        int index = 0;
        for(int i = 0; i < nums.length; i++){
            if(nums[i] > current){
                current = nums[i];
                nums[index] = current;
                index++;
            }
        }
        return index;
    }
}