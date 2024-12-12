/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
void backtrack(int* nums, int numsSize, int** result, int* returnSize, int* tempList, int* used, int** returnColumnSizes, int depth) {
    if (depth == numsSize) {
        result[*returnSize] = (int*)malloc(numsSize * sizeof(int));
        memcpy(result[*returnSize], tempList, numsSize * sizeof(int));
        (*returnColumnSizes)[*returnSize] = numsSize;
        (*returnSize)++;
        return;
    }
    for (int i = 0; i < numsSize; i++) {
        if (used[i] || (i > 0 && nums[i] == nums[i - 1] && !used[i - 1])) continue;
        used[i] = 1;
        tempList[depth] = nums[i];
        backtrack(nums, numsSize, result, returnSize, tempList, used, returnColumnSizes, depth + 1);
        used[i] = 0;
    }
}

int compare(const void* a, const void* b) {
    return (*(int*)a - *(int*)b);
}

int** permuteUnique(int* nums, int numsSize, int* returnSize, int** returnColumnSizes) {
    qsort(nums, numsSize, sizeof(int), compare);
    int perms = 1;
    for (int i = 2; i <= numsSize; i++) perms *= i;
    int** result = (int**)malloc(perms * sizeof(int*));
    *returnColumnSizes = (int*)malloc(perms * sizeof(int));
    *returnSize = 0;
    int* tempList = (int*)malloc(numsSize * sizeof(int));
    int* used = (int*)calloc(numsSize, sizeof(int));
    backtrack(nums, numsSize, result, returnSize, tempList, used, returnColumnSizes, 0);
    free(tempList);
    free(used);
    return result;
}
