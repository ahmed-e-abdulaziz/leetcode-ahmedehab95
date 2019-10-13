/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        if(head.next == null && n == 1)
            return null;
        remove(head, 0, n);
        return head; 
    }
    
    public int remove(ListNode node, int index, int n) {
        if(node.next == null){
            if(n == 1){
                return -1;
            }
            return 1;
        }
        index = remove(node.next, index, n);
        if(index == -1){
            node.next = null;
            return 2;
        }
        index++;
        if(index == n){
            node.val = node.next.val;
            node.next = node.next.next;
        }
        return index;
    }
}