#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

int cmpItems(const void* a, const void* b) {
    int* x = *(int**)a;
    int* y = *(int**)b;
    return x[0] - y[0];
}

int cmpQueries(const void* a, const void* b) {
    int* x = (int*)a;
    int* y = (int*)b;
    return x[0] - y[0];
}

int* maximumBeauty(int** items, int itemsSize, int* itemsColSize, int* queries, int queriesSize, int* returnSize) {
    *returnSize = queriesSize;
    int* res = (int*)malloc(sizeof(int) * queriesSize);
    qsort(items, itemsSize, sizeof(int*), cmpItems);
    int** qrs = (int**)malloc(queriesSize * sizeof(int*));
    for (int i = 0; i < queriesSize; i++) {
        qrs[i] = (int*)malloc(2 * sizeof(int));
        qrs[i][0] = queries[i];
        qrs[i][1] = i;
    }
    qsort(qrs, queriesSize, sizeof(int*), cmpQueries);
    int* maxB = (int*)malloc(sizeof(int) * itemsSize);
    maxB[0] = items[0][1];
    for (int i = 1; i < itemsSize; i++) {
        maxB[i] = items[i][1] > maxB[i - 1] ? items[i][1] : maxB[i - 1];
    }
    for (int i = 0, j = 0; i < queriesSize; i++) {
        int q = qrs[i][0];
        int idx = qrs[i][1];
        int l = 0, r = itemsSize - 1, pos = -1;
        while (l <= r) {
            int m = l + (r - l) / 2;
            if (items[m][0] <= q) {
                pos = m;
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
        res[idx] = pos == -1 ? 0 : maxB[pos];
    }
    for (int i = 0; i < queriesSize; i++) {
        free(qrs[i]);
    }
    free(qrs);
    free(maxB);
    return res;
}
