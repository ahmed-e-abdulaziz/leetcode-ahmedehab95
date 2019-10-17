class Solution {
    int[] cache = new int[50];
    {cache[0]=1;
    cache[1]=1;}
    public int climbStairs(int n) {
        if(cache[n] > 0)
            return cache[n];
        else
            cache[n] = climbStairs(n-1) + climbStairs(n-2);
        return cache[n];
    }
}