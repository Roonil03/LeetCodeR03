/**
 * Definition for singly-linked list.
 * */
#include <stdio.h>
#include <stdlib.h>

 struct ListNode {
     int val;
     struct ListNode *next;
 };
 
typedef struct ListNode Node;

void insertNode(Node** head, int v) {
    Node* newNode = (Node*)malloc(sizeof(Node));
    newNode->val = v;
    newNode->next = NULL;
    if (*head == NULL) {
        *head = newNode;
    } else {
        Node* n1 = *head;
        while (n1->next != NULL) {
            n1 = n1->next;
        }
        n1->next = newNode;
    }
}

struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2) {
    Node* n1 = list1;
    Node* n2 = list2;
    Node* result = NULL;
    while (n1 != NULL && n2 != NULL) {
        if (n1->val < n2->val) {
            insertNode(&result, n1->val);
            n1 = n1->next;
        } else {
            insertNode(&result, n2->val);
            n2 = n2->next;
        }
    }
    while (n1 != NULL) {
        insertNode(&result, n1->val);
        n1 = n1->next;
    }
    while (n2 != NULL) {
        insertNode(&result, n2->val);
        n2 = n2->next;
    }
    return result;
}
