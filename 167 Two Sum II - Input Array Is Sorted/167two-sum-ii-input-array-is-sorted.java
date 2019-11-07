class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int low = 0;
        int high = numbers.length -1;
        while(high > low){
            int val = numbers[high] + numbers[low];
            if(val == target) return new int[]{low+1, high+1};
            else if(val > target) high--;
                else if(val < target) low++;
        }
        return null;
    }
}