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
    bool isPalindrome(ListNode* head) {
        ListNode* s = head, *f = head;
        while(f->next && f->next->next){
            s = s->next;
            f = f->next->next;
        }
        ListNode* p = nullptr, *c = s->next;
        while(c) {
            ListNode *n = c->next;
            c->next = p;
            p = c;
            c = n;
        }
        ListNode* fr = head, *sc = p;
        while(sc){
            if(fr->val != sc->val){
                return false;
            }
            fr = fr->next;
            sc = sc->next;
        }
        return true;
    }
};