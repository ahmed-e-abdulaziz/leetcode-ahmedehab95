/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public boolean isPalindrome(ListNode head) {
        
        if(head == null || head.next == null)
            return true;
        
        ListNode tmp = head, slow = head, fast = head;
        
		// Finding the node at the middle.
		// We will split the list from middle point.
        while(fast != null && fast.next != null){
            fast = fast.next.next;
            tmp = slow;
            slow = slow.next;
        }
        
        ListNode revH = tmp.next;
        tmp.next = null;
        
		// reversing the second half of the list.
        revH = rev(revH);
        
        while(head != null){
            if(head.val != revH.val)
                return false;
            head = head.next;
            revH = revH.next;
        }
        
        return true;
    }
    
    private static ListNode rev(ListNode node){
        if(node == null || node.next == null)
            return node;
        ListNode a = node, b = node.next;
        
        while(b != null){
            ListNode tmp = b.next;
            b.next = a;
            a = b;
            b = tmp;
        }
        
        return a;
    }
}