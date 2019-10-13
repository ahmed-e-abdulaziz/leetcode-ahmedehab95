/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if(l1 == null)
            return l2;
        if(l2 == null)
            return l1;
        ListNode head = new ListNode(l1.val < l2.val? l1.val:l2.val);
        //head.next = new ListNode(l1.val < l2.val? l2.val: l1.val);
        if(l1.val < l2.val)
            l1 = l1.next;
        else
            l2 = l2.next;
        ListNode node = head;
        while(l1 != null || l2 != null){
            if(l1 != null && l2 != null){
                if(l1.val < l2.val){
                    node.next = new ListNode(l1.val);
                    node = node.next;
                    l1 = l1.next;
                } else {
                    node.next = new ListNode(l2.val);
                    node = node.next;
                    l2 = l2.next;
                }
            }else{
                if(l1 != null){
                    node.next = new ListNode(l1.val);
                    node = node.next;
                    l1 = l1.next;
                }
                if(l2 != null){
                    node.next = new ListNode(l2.val);
                    node = node.next;
                    l2 = l2.next;
                }
            }
        }
        return head;
    }
}