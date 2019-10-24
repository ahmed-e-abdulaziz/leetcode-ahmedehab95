/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int len1 = 0;
        int len2 = 0;
        ListNode temp1 = l1;
        ListNode temp2 = l2;
        while(l1!= null)
        {
            len1++;
            l1 = l1.next;
        }
        while(l2 != null)
        {
            len2++;
            l2 = l2.next;
        }
        l1 = temp1;
        l2 = temp2;
        if(len1 > len2)
        {
            temp1 = l2;
            temp2 = l1; 
        }
         add(temp1,temp2);
        return process(temp2); 
    }
    public void add(ListNode l1,ListNode l2)
    {
        while(l1 != null && l2 != null)
        {
            l2.val += l1.val;
            l1 = l1.next;
            l2 = l2.next;
        }
    }
    public ListNode process(ListNode temp)
    {
        ListNode l = temp;
        while(l != null)
        {
            if(l.val < 10)
                l = l.next;
            else 
            {
                l.val -=10;
               if(l.next != null)
               {
                   l.next.val += 1;
                   l = l.next;
                }else
                    l.next = new ListNode(1);
               }
        }
        return temp;
    }
}