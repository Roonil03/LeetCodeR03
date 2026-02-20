/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;
};
*/

class Solution {
public:
    Node* flatten(Node* head) {
        if(!head){
            return nullptr;
        }
        Node* cur = head;
        while(cur){
            if(cur->child){
                Node* a = cur->child;
                while(a->next){
                    a = a->next;
                }
                a->next = cur->next;
                if(cur->next){
                    cur->next->prev = a;
                }
                cur->next = cur->child;
                cur->child->prev = cur;
                cur->child = nullptr;
            }
            cur = cur->next;
        }
        return head;
    }
};