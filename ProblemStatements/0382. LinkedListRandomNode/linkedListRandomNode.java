public /**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {

    ListNode head;
    Random rand;

    public Solution(ListNode head) {
        this.head = head;
        rand = new Random();
    }
    
    public int getRandom() {
        ListNode cur = head;
        int count = 0, res = 0;
        while(cur != null){
            count++;
            if(rand.nextInt(count) == 0){
                res = cur.val;
            }
            cur = cur.next;
        }
        return res;
    }
}

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(head);
 * int param_1 = obj.getRandom();
 */ {
    
}
