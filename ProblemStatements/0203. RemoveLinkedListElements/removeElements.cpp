/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode d(0);
        d.next = head;
        ListNode* cur = &d;
        while(cur->next){
            if(cur->next->val == val){
                ListNode* del = cur->next;
                cur->next = cur->next->next;
                delete del;
            } else{
                cur = cur->next;
            }
        }
        return d.next;
    }
};