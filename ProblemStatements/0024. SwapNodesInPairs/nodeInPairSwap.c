#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;  
    struct ListNode *next;
};

typedef struct ListNode Node;

void swap(Node* head)
{
    if(head == NULL) return;
    Node* temp1 = head;
    Node* temp2 = head->next;
    if(temp2 == NULL) return;
    int t = temp1->val;
    temp1->val = temp2->val;
    temp2->val = t;
    swap(temp2->next);
}
struct ListNode* swapPairs(struct ListNode* head) {
    swap(head);
    return head;
}