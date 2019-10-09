class Solution {
    public void reverseString(char[] s) {
        int j = s.length - 1;
        int i = 0;
        while(j > i) {
            char temp = s[i];
            s[i] = s[j];
            s[j] = temp;
            j--;
            i++;
        }
    }
}