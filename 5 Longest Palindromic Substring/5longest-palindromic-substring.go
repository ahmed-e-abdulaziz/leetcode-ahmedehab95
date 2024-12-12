func longestPalindrome(s string) string {
    
    var ans string = ""
    
    for i := 0; i< len(s); i++ {
        step:= 1
        current := string(s[i])
        noCenter := true
        center := true
        noCenCur := ""

        for noCenter || center {
            
            left := i - step
            right := i + step
            noCenterLeft := i - step + 1
            noCenterRight := i + step 
            
            //Center Case 
            if left > -1 && right < len(s) && center  {
                if s[left] == s[right] {
                    current = s[left:right+1]
                } else {
                    center = false
                }
            } else {
                center = false
            }
            
            // No Center Case 
            if noCenterLeft > -1 && noCenterRight < len(s) && noCenter {
                 if s[noCenterLeft] == s[noCenterRight] {
                    noCenCur = s[noCenterLeft:noCenterRight+1]
                } else {
                    noCenter = false
                }
            }else {
                noCenter = false
            }
            step++
        }
        
        if len(current )> len(ans) { ans = current }
        if len(noCenCur) > len(ans) { ans = noCenCur}
    }
    return ans
}