#include <stdbool.h>

#define MAX 10000

typedef struct {
    int data[MAX];
    bool exists[MAX];
} HashSet;

void initHashSet(HashSet *set) {
    for (int i = 0; i < MAX; i++) {
        set->exists[i] = false;
    }
}

void insertHashSet(HashSet *set, int value) {
    int idx = abs(value) % MAX;
    set->data[idx] = value;
    set->exists[idx] = true;
}

bool containsHashSet(HashSet *set, int value) {
    int idx = abs(value) % MAX;
    return set->exists[idx] && set->data[idx] == value;
}

bool checkIfExist(int* arr, int arrSize) {
    HashSet set;
    initHashSet(&set);

    for (int i = 0; i < arrSize; i++) {
        int num = arr[i];
        if (containsHashSet(&set, 2 * num) || (num % 2 == 0 && containsHashSet(&set, num / 2))) {
            return true;
        }
        insertHashSet(&set, num);
    }
    return false;
}