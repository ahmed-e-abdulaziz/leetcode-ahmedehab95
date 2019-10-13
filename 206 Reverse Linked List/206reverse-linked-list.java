/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseList(ListNode head) {
        if(head == null || head.next == null)
            return head;
        return reverse(head);
    }
    
    public ListNode reverse(ListNode node){
        if(node.next == null)
            return node;
        ListNode p = reverse(node.next);
        node.next.next = node;
        node.next = null;
        return p;
    }
}