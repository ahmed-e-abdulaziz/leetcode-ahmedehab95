class Solution {
    public int firstUniqChar(String s) {
        if(s == null || s.equals("")){
            return -1;
        }
        int[] lettersCount = new int[26];
        int[] orderArray = new int[26];
        int order = 0;
        for(char c : s.toCharArray()){
            int pos = (int) c - 'a';
            lettersCount[pos]++;
            if(orderArray[pos] == 0)
                orderArray[pos] = order;
            order++;
        }
        int leastOrder = Integer.MAX_VALUE;
        for(int i = 0; i < 26; i++) {
             if(lettersCount[i] == 1)
                 leastOrder = orderArray[i] < leastOrder ? orderArray[i] : leastOrder ;
        }
        if(leastOrder == Integer.MAX_VALUE){
            return -1;
        }
        return leastOrder;
    }
}