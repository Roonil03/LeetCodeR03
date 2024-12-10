#include <stdlib.h>

#define MAX 10007

typedef struct Hash {
    char *substr;
    int count;
    struct Hash *next;
} Hash;

unsigned long hashing(const char *str, int len) {
    unsigned long hash = 5381;
    for (int i = 0; i < len; i++) {
        hash = ((hash << 5) + hash) + str[i];
    }
    return hash % MAX;
}

Hash *createNode(const char *substr, int len) {
    Hash *node = (Hash *)malloc(sizeof(Hash));
    node->substr = (char *)malloc(len + 1);
    strncpy(node->substr, substr, len);
    node->substr[len] = '\0';
    node->count = 1;
    node->next = NULL;
    return node;
}

void insertSubstring(Hash **h1, const char *substr, int len) {
    unsigned long hash = hashing(substr, len);
    Hash *h = h1[hash];
    while (h) {
        if (strncmp(h->substr, substr, len) == 0) {
            h->count++;
            return;
        }
        h = h->next;
    }
    Hash *newNode = createNode(substr, len);
    newNode->next = h1[hash];
    h1[hash] = newNode;
}

int getSubstringCount(Hash **h1, const char *substr, int len) {
    unsigned long hash = hashing(substr, len);
    Hash *h = h1[hash];
    while (h) {
        if (strncmp(h->substr, substr, len) == 0) {
            return h->count;
        }
        h = h->next;
    }
    return 0;
}

void freeHashTable(Hash **h1) {
    for (int i = 0; i < MAX; i++) {
        Hash *h = h1[i];
        while (h) {
            Hash *temp = h;
            h = h->next;
            free(temp->substr);
            free(temp);
        }
    }
}

int maximumLength(char *s) {
    int len = strlen(s), maxlen = -1;
    for (int sub = 1; sub <= len; sub++) {
        Hash *h1[MAX] = {NULL};
        for (int i = 0; i <= len - sub; i++) {
            int valid = 1;
            for (int j = i; j < i + sub - 1; j++) {
                if (s[j] != s[j + 1]) {
                    valid = 0;
                    break;
                }
            }
            if (valid) {
                insertSubstring(h1, s + i, sub);
            }
        }
        for (int i = 0; i <= len - sub; i++) {
            if (getSubstringCount(h1, s + i, sub) >= 3) {
                maxlen = sub > maxlen ? sub : maxlen;
            }
        }
        freeHashTable(h1);
    }
    return maxlen == -1 ? -1 : maxlen;
}
