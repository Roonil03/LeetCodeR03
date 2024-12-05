#include <stdlib.h>
/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
void backtrack(int* candidates, int candidatesSize, int target, int start, int* temp, int tempSize, int** result, int* returnSize, int** returnColumnSizes) {
    if (target == 0) {
        result[*returnSize] = malloc(tempSize * sizeof(int));
        for (int i = 0; i < tempSize; i++) {
            result[*returnSize][i] = temp[i];
        }
        (*returnColumnSizes)[*returnSize] = tempSize;
        (*returnSize)++;
        return;
    }
    for (int i = start; i < candidatesSize; i++) {
        if (i > start && candidates[i] == candidates[i - 1]) continue;
        if (candidates[i] > target) break;
        temp[tempSize] = candidates[i];
        backtrack(candidates, candidatesSize, target - candidates[i], i + 1, temp, tempSize + 1, result, returnSize, returnColumnSizes);
    }
}

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int** combinationSum2(int* candidates, int candidatesSize, int target, int* returnSize, int** returnColumnSizes) {
    *returnSize = 0;
    int** result = malloc(1000 * sizeof(int*));
    *returnColumnSizes = malloc(1000 * sizeof(int));
    int* temp = malloc(candidatesSize * sizeof(int));
    qsort(candidates, candidatesSize, sizeof(int), cmp);
    backtrack(candidates, candidatesSize, target, 0, temp, 0, result, returnSize, returnColumnSizes);
    free(temp);
    return result;
}
