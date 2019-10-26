/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode middleNode(ListNode head) {
        ListNode middle = head;
        while(head != null && head.next != null){
            head = head.next.next;
            middle = middle.next;
        }
        return middle;
    }
}