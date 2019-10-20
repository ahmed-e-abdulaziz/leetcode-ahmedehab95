class Solution {
    public List<String> fizzBuzz(int n) {
        ArrayList<String> strlist = new ArrayList<>();
        for(int i = 1; i <= n; i++)
        {
            if(i % 5 == 0 && i % 3 == 0)
            {
                strlist.add("FizzBuzz");
            } else {
                if(i % 5 == 0)
                    strlist.add("Buzz");
                else {
                if(i % 3 == 0)
                    strlist.add("Fizz");
                else strlist.add(Integer.toString(i));
                }
            }
        }
        return strlist;
    }
}