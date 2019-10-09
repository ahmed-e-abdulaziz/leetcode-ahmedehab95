class Solution {
    public int reverse(int x) {
        try{
            StringBuilder reversed = new StringBuilder(Integer.toString(x)).reverse();
            if(reversed.charAt(reversed.length()-1) == '-'){
                reversed.deleteCharAt(reversed.length()-1);
                reversed.insert(0,"-");
            }
                return Integer.valueOf(reversed.toString()); 
        }
        catch(Exception ex){
            return 0;
        }
    }
}