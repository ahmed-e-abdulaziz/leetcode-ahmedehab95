class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        if(nums1.length == 0){
            nums1 = nums2;
            return;
        }
        
        if(nums2.length == 0)
            return;
        for(int i = m-1; i >=0 ; i--){
            nums1[i+n]=nums1[i];
        }
        int i = n;
        int j = 0;
        int k = 0;
        while(i < nums1.length || j < n)
        {
            if(i == nums1.length){
                nums1[k] = nums2[j++];
            } else {
                if(j == nums2.length){
                    nums1[k] = nums1[i++];
                } else {
                    if(nums1[i] < nums2[j])
                    {
                        nums1[k] = nums1[i];
                        i++;
                    } else {
                        nums1[k] = nums2[j];
                        j++;
                    }
               }
            }
            k++;
        }
    }
}