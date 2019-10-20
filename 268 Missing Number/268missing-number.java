class Solution {
    public int missingNumber(int[] nums) {
        if(nums.length==1){
            if(nums[0] == 0)
                return 1;
            if(nums[0] == 1)
                return 0;
        }
        
        int sum = 0;
        int max = -1;
        boolean zeroOccured = false;
        for(int i = 0; i < nums.length; i++){
            if(max < nums[i])
                max = nums[i];
            if(nums[i]==0)
                zeroOccured = true;
            sum+=nums[i];
        }
        for(int i = 0; i< max+1; i++)
            sum-=i;
        if(sum == 0 && zeroOccured)
            return max+1;
        else if(!zeroOccured) return 0;
        return sum*-1;
    }
}