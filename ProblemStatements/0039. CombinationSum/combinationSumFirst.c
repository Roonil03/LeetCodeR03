#include <stdlib.h>
/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */

typedef struct {
    int** data;
    int* colSizes;
    int size;
    int capacity;
} Result;

void addCombination(Result* result, int* combination, int combSize) {
    if (result->size == result->capacity) {
        result->capacity *= 2;
        result->data = realloc(result->data, result->capacity * sizeof(int*));
        result->colSizes = realloc(result->colSizes, result->capacity * sizeof(int));
    }
    result->data[result->size] = malloc(combSize * sizeof(int));
    for (int i = 0; i < combSize; i++) {
        result->data[result->size][i] = combination[i];
    }
    result->colSizes[result->size] = combSize;
    result->size++;
}

void backtrack(int* candidates, int candidatesSize, int target, int start, int* combination, int combSize, Result* result) {
    if (target == 0) {
        addCombination(result, combination, combSize);
        return;
    }
    for (int i = start; i < candidatesSize; i++) {
        if (candidates[i] > target) continue;
        combination[combSize] = candidates[i];
        backtrack(candidates, candidatesSize, target - candidates[i], i, combination, combSize + 1, result);
    }
}

int** combinationSum(int* candidates, int candidatesSize, int target, int* returnSize, int** returnColumnSizes) {
    Result result = { malloc(16 * sizeof(int*)), malloc(16 * sizeof(int)), 0, 16 };
    int* combination = malloc(target * sizeof(int));
    backtrack(candidates, candidatesSize, target, 0, combination, 0, &result);
    free(combination);
    *returnSize = result.size;
    *returnColumnSizes = result.colSizes;
    return result.data;
}
