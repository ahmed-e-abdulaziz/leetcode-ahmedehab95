class Solution {
    public int maxDistance(int[][] grid) {
        int[][] dist = new int[grid.length][grid[0].length];
        for(int i = 0; i < dist.length; i++){
            Arrays.fill(dist[i], Integer.MAX_VALUE);
        }
        for(int j = 0; j < grid.length; j++){
            for(int i = 0; i < grid[0].length; i++)
            {
                if(grid[j][i] == 1)
                {
                    for(int l = 0; l < grid.length; l++){
                        for(int k = 0; k < grid[0].length; k++)
                        {
                            dist[l][k] = Math.min(dist[l][k], Math.abs(l-j)+Math.abs(k-i));
                        }
                    }
                }
            }
        }
        int max = -1;
        for(int[] row : dist)
            for(int cell : row)
                max = Math.max(cell,max);
        
        return max == Integer.MAX_VALUE || max == 0 ? -1 : max;
    }
}