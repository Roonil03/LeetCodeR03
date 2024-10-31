#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

typedef struct {
    char *data;
    int top;
    int capacity;
} Stack;

Stack* createStack(int capacity) {
    Stack* s = (Stack*)malloc(sizeof(Stack));
    s->capacity = capacity;
    s->top = -1;
    s->data = (char*)malloc(capacity * sizeof(char));
    return s;
}

bool isEmpty(Stack* s) {
    return s->top == -1;
}

void push(Stack* s, char value) {
    s->data[++s->top] = value;
}

char pop(Stack* s) {
    return s->data[s->top--];
}

char peek(Stack* s) {
    return s->data[s->top];
}

bool isMatchingPair(char open, char close) {
    return (open == '(' && close == ')') || (open == '{' && close == '}') || (open == '[' && close == ']');
}

bool isValid(char* str) {
    int n = strlen(str);
    Stack* s = createStack(n);    
    for (int i = 0; i < n; i++) {
        char current = str[i];
        if (current == '(' || current == '{' || current == '[') {
            push(s, current);
        } else {
            if (isEmpty(s) || !isMatchingPair(pop(s), current)) {
                free(s->data);
                free(s);
                return false;
            }
        }
    }
    bool result = isEmpty(s);
    free(s->data);
    free(s);
    return result;
}