class Solution {

    private int[] arr;
    private Random generator;
    
    public Solution(int[] nums) {
        this.arr = new int[nums.length];
        this.generator = new Random(System.nanoTime());
        for (int i = 0; i < nums.length; i++) this.arr[i] = nums[i];
    }
    
    /** Resets the array to its original configuration and return it. */
    public int[] reset() {
        return arr;
    }
    
    /** Returns a random shuffling of the array. */
    public int[] shuffle() {
        int[] copy = new int[arr.length];

        int ind = 0;
        for (int i = 0; i < arr.length; i++) {
            while ((ind = generator.nextInt(arr.length)) < 0 || copy[ind] != 0) continue;
            copy[ind] = arr[i];
        }
        
        return copy;
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(nums);
 * int[] param_1 = obj.reset();
 * int[] param_2 = obj.shuffle();
 */