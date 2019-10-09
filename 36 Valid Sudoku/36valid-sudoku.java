class Solution {
    public boolean isValidSudoku(char[][] board) {
        for(int i = 0; i < 9; i++){
            boolean[] exists = new boolean[9];
            for(int j = 0; j < 9; j++){
                if(board[i][j] != '.'){
                    if(exists[Character.getNumericValue(board[i][j])-1]){
                        return false;
                    }
                    exists[Character.getNumericValue(board[i][j])-1] = true;
                }
            }
        }
        
        for(int i = 0; i < 9; i++){
            boolean[] exists = new boolean[9];
            for(int j = 0; j < 9; j++){
                if(board[j][i] != '.'){
                    if(exists[Character.getNumericValue(board[j][i])-1]){
                        return false;
                    }
                    exists[Character.getNumericValue(board[j][i])-1] = true;
                }
            }
        }
        
        for(int i = 0; i < 9; i+=3){
            for(int j = 0; j < 9; j+=3){
                boolean[] exists = new boolean[9];
                for(int k = i; k < i+3; k++){    
                    for(int l = j; l < j+3; l++){
                        if(board[k][l] != '.'){
                            if(exists[Character.getNumericValue(board[k][l])-1]){
                                return false;
                            }
                            exists[Character.getNumericValue(board[k][l])-1] = true;
                        }
                    }
                }
            }
        }
        
        
        return true;
    }
}