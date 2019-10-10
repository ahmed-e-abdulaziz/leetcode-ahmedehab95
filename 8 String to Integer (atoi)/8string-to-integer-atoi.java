class Solution {
    public int myAtoi(String str) {
        str = str.trim();
        int result = 0;
        int modifier = 1;
        if(str.length() == 0 || (str.charAt(0) == '-' || str.charAt(0) == '+') && str.length() == 1 ){
            return 0;
        }
        int i = 1;
        if(!((str.charAt(0) >= '0' && str.charAt(0) <= '9') || ((str.charAt(0) == '-' || str.charAt(0) == '+')  && str.charAt(1) >= '0' && str.charAt(1) <= '9'))){
           return 0;
        } else {
            if(str.charAt(0) == '-') {
                result = Character.getNumericValue(str.charAt(1))*-1;
                modifier = -1;
                i++;
            } else {
                if(str.charAt(0) == '+'){
                    result = Character.getNumericValue(str.charAt(1));
                    i++;
                } else {
                    result = Character.getNumericValue(str.charAt(0));
                }
            }
        }
        
        while(i < str.length()){
            if(!(str.charAt(i) >= '0' && str.charAt(i) <= '9')){
                break;
            }
            
            if(i >= 9){
                long possibleOverflow = ((long)result)*10+ (long)Character.getNumericValue(str.charAt(i))*modifier;
                System.out.println(possibleOverflow);
                switch(modifier){
                    case 1: 
                        if(possibleOverflow>=Integer.MAX_VALUE){
                            return Integer.MAX_VALUE;
                        }
                        break;
                    case -1:
                        if(possibleOverflow<=Integer.MIN_VALUE){
                            return Integer.MIN_VALUE;
                        }
                        break;
                }        
            }

            result = (int) (result * 10 )+ Character.getNumericValue(str.charAt(i))*modifier;
            i++;
        }
        return result;
    }
    
}