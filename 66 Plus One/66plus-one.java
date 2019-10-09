class Solution {
    public int[] plusOne(int[] digits) {
        if(digits[digits.length-1] == 9){
            return dealWithNines(digits);
        } else {
            digits[digits.length-1] = digits[digits.length-1]+1;
            return digits;
        }
    }
    
    private int[] dealWithNines(int[] digits){
        int i = digits.length-1;
        while(i >= 0 && digits[i] == 9){
            i--;
        }
        System.out.println(i);
        if(i == -1){
           int[] resultArray = new int[digits.length+1];
            resultArray[0] = 1;
            for(int x = 1; x < digits.length-1; x++){
                resultArray[x] = 0;
            }
            return resultArray;
        } else {
            digits[i] = digits[i] + 1;
            i++;
            while(i < digits.length){
                digits[i] = 0;
                i++;
            }
            return digits;
        }
    }
}