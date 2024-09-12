#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct ListNode {
    int val;
    struct ListNode *next;
} Node;

void createNode(int val, Node** s) {
    *s = (Node*)malloc(sizeof(Node));
    (*s)->val = val;
    (*s)->next = NULL;
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    struct ListNode* result = (struct ListNode*)malloc(sizeof(struct ListNode));
    struct ListNode* head = result;
    struct ListNode* end;
    int carry = 0;
    while(l1 != NULL || l2 != NULL){
        result->next = NULL;
        if(l1 == NULL && l2 != NULL){
            result->val = (carry + l2->val)%10;
            carry = (carry + l2->val)/10;
            l2 = l2->next;
        }
        else if(l2 == NULL && l1 != NULL){
            result->val = (carry + l1->val)%10;
            carry = (carry + l1->val)/10;
            l1 = l1->next;
        }
        else{
            result->val = (carry + l1->val + l2->val)%10;
            carry = (carry + l1->val + l2->val)/10;
            l2 = l2->next;
            l1 = l1->next;
        }
        result->next = (struct ListNode*)malloc(sizeof(struct ListNode));
        end = result;
        result = result->next;
    }
    if(carry != 0){
    result->val = carry;
    result->next = NULL;
    }
    else{
         end->next = NULL;
    }
    return head;
}

int* reverse(char* str) {
    int len = strlen(str);
    int* n = (int*)malloc(sizeof(int) * len);
    for (int i = 0; i < len; i++) {
        n[i] = str[len - 1 - i] - '0';
    }
    return n;
}

Node* intoList(int* n, int len) {
    Node* head = NULL;
    Node* tail = NULL;

    for (int i = 0; i < len; i++) {
        Node* newNode;
        createNode(n[i], &newNode);

        if (head == NULL) {
            head = newNode;
        } else {
            tail->next = newNode;
        }

        tail = newNode;
    }

    return head;
}

void freeList(Node* head) {
    Node* current = head;
    Node* next;
    while (current != NULL) {
        next = current->next;
        free(current);
        current = next;
    }
}

int main() {
    char str1[50], str2[50];
    printf("Enter the two numbers:\n");
    scanf("%s", str1);
    scanf("%s", str2);

    int* n1 = reverse(str1);
    int* n2 = reverse(str2);

    int len1 = strlen(str1);
    int len2 = strlen(str2);

    Node* l1 = intoList(n1, len1);
    Node* l2 = intoList(n2, len2);

    Node* ans = addTwoNumbers(l1, l2);

    while (ans != NULL) {
        printf("%d ", ans->val);
        ans = ans->next;
    }
    printf("\n");

    free(n1);
    free(n2);
    freeList(l1);
    freeList(l2);
    freeList(ans);

    return EXIT_SUCCESS;
}