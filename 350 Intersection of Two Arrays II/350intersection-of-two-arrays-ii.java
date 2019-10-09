import java.util.*;
class Solution {
        public int[] intersect(int[] nums1, int[] nums2) {
        boolean oneBigger = nums1.length > nums2.length;
        ArrayList<Integer> result = new ArrayList<>();
        if(oneBigger){
            HashMap<Integer, Integer> map = new HashMap<>();
            for(int i = 0; i < nums2.length; i++){
                Integer count = 0;
                if(map.containsKey(new Integer(nums2[i]))){
                    int oldCount = map.get(new Integer(nums2[i]));
                    oldCount++;
                    count = new Integer(oldCount);
                } else { 
                    count = 1;
                }
                map.put(new Integer(nums2[i]), count);
            }
            for(int i = 0; i < nums1.length; i++){
                if(map.containsKey(new Integer(nums1[i])) && map.get(new Integer(nums1[i])) > 0){
                    result.add(new Integer(nums1[i]));
                    map.put(new Integer(nums1[i]),map.get(new Integer(nums1[i]))-1);
                }
            }
        } else {
            HashMap<Integer, Integer> map = new HashMap<>();
            for(int i = 0; i < nums1.length; i++){
                Integer count = 0;
                if(map.containsKey(new Integer(nums1[i]))){
                    int oldCount = map.get(new Integer(nums1[i]));
                    oldCount++;
                    count = new Integer(oldCount);
                } else { 
                    count = 1;
                }
                map.put(new Integer(nums1[i]), count);
            }
            for(int i = 0; i < nums2.length; i++){
                if(map.containsKey(new Integer(nums2[i])) && map.get(new Integer(nums2[i])) > 0){
                    result.add(new Integer(nums2[i]));
                    map.put(new Integer(nums2[i]),map.get(new Integer(nums2[i]))-1);
                }
            }
        }
        return result.stream().mapToInt(i -> i).toArray();
    }
}