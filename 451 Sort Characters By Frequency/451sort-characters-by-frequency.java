class Solution {
    public String frequencySort(String s) {
        int[] frq = new int[256];
        HashMap<Character, Integer> map = new HashMap<>();
        for(int i= 0; i<s.length(); i++)
            map.put(s.charAt(i), map.containsKey(s.charAt(i))?map.get(s.charAt(i))+1:1);
        StringBuilder sb = new StringBuilder();
        for(Map.Entry<Character, Integer> entry : map.entrySet().stream().sorted((e1,e2) -> e2.getValue().compareTo(e1.getValue())).collect(Collectors.toList())){
            for(int i = 0; i < entry.getValue(); i++) {
                sb.append(entry.getKey());
            }
        }
        return sb.toString();
    }
}