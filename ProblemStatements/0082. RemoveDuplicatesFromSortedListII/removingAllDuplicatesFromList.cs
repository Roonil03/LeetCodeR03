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
    public ListNode DeleteDuplicates(ListNode head) {
        try{
            ListNode n = new ListNode(0, null);
            n.next = head;
            ListNode prev = n, cur = head;
            while(cur != null)
            {
                if(cur.next != null && cur.val == cur.next.val){
                    while(cur.next != null && cur.val == cur.next.val){
                        cur = cur.next;
                    }
                    prev.next = cur.next;
                }
                else{
                    prev = prev.next;
                }
                cur = cur.next;
            }
            return n.next;
        }
        catch
        {
            Console.WriteLine("error");
            return null;
        }
    }
}