#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

#define HASH_SIZE 1000

typedef struct Node {
    char* key;
    struct Node* next;
} Node;

typedef struct {
    Node* buckets[HASH_SIZE];
} HashMap;

unsigned int hash(const char* str) {
    unsigned int hash = 0;
    while (*str) {
        hash = (hash * 31) + *str;
        str++;
    }
    return hash % HASH_SIZE;
}

HashMap* createHashMap() {
    HashMap* map = (HashMap*)malloc(sizeof(HashMap));
    for (int i = 0; i < HASH_SIZE; i++) {
        map->buckets[i] = NULL;
    }
    return map;
}

bool containsKey(HashMap* map, const char* key) {
    unsigned int index = hash(key);
    Node* node = map->buckets[index];
    while (node != NULL) {
        if (strcmp(node->key, key) == 0) {
            return true;
        }
        node = node->next;
    }
    return false;
}

void insertKey(HashMap* map, const char* key) {
    unsigned int index = hash(key);
    Node* node = (Node*)malloc(sizeof(Node));
    node->key = strdup(key);
    node->next = map->buckets[index];
    map->buckets[index] = node;
}

void removeKey(HashMap* map, const char* key) {
    unsigned int index = hash(key);
    Node* node = map->buckets[index];
    Node* prev = NULL;
    while (node != NULL) {
        if (strcmp(node->key, key) == 0) {
            if (prev == NULL) {
                map->buckets[index] = node->next;
            } else {
                prev->next = node->next;
            }
            free(node->key);
            free(node);
            return;
        }
        prev = node;
        node = node->next;
    }
}

void dfs(char* s, int start, int len, HashMap* used, int* maxCount, int count) {
    if (start == len) {
        if (count > *maxCount) {
            *maxCount = count;
        }
        return;
    }

    char temp[51];
    for (int i = start; i < len; i++) {
        int subLen = i - start + 1;
        strncpy(temp, s + start, subLen);
        temp[subLen] = '\0';

        if (!containsKey(used, temp)) {
            insertKey(used, temp);
            dfs(s, i + 1, len, used, maxCount, count + 1);
            removeKey(used, temp); 
            }
    }
}

int maxUniqueSplit(char* s) {
    int maxCount = 0;
    HashMap* used = createHashMap();
    dfs(s, 0, strlen(s), used, &maxCount, 0);
    //freeHashMap(used);
    return maxCount;
}