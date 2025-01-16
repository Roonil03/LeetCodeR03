/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode Partition(ListNode head, int x) {
       ListNode l = new ListNode(0), r = new ListNode(0);
       ListNode a = l, b = r;
       while(head!=null)
       {
            if(head.val < x)
            {
                a.next = head;
                a = a.next;
            }
            else{
                b.next = head;
                b = b.next;
            }
            head = head.next;
       }
       a.next = r.next;
       b.next = null;
       return l.next;
    }
}