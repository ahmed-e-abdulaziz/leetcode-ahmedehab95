class Solution {
    public int strStr(String haystack, String needle) {
        int nlength = needle.length();
        int hlength = haystack.length();
        if(nlength == 0)
            return 0;
        if(nlength == 1)
            return haystack.indexOf(needle.charAt(0));
        if(hlength < nlength)
            return -1;
        char[] hsarr = haystack.toCharArray();
        char[] narr = needle.toCharArray();
        for(int i = 0; i < hlength; i++){
            if(hsarr[i] == narr[0]){
                int j = 1;
                while(j < nlength){
                    if(i+j >= hlength)
                        return -1;
                    if(hsarr[i+j] == narr[j]){
                        if(j == nlength-1)
                            return i;
                        j++;
                    } else {break;}
                }
            }
        }
        return -1;
    }
}