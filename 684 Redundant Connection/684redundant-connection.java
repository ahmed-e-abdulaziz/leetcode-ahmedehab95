class Solution {
    int[] parent;
    public int[] findRedundantConnection(int[][] edges) {
        parent = new int[edges.length+1];
        for(int i = 0; i < edges.length ; i++){
            int a = findParent(edges[i][0]);
            int b = findParent(edges[i][1]);
            if(a == b)
                return edges[i];
            parent[b] = a;
        }
        return new int[2];
    }

    private int findParent(int index) {
            if(parent[index] == 0)
                return index;
            return findParent(parent[index]);
    }
}