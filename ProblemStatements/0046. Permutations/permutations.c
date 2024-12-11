/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

void p(int *nums, int start, int end, int **result, int *returnSize, int *columnSizes) {
    if (start == end) {
        result[*returnSize] = malloc((end + 1) * sizeof(int));
        for (int i = 0; i <= end; i++) {
            result[*returnSize][i] = nums[i];
        }
        columnSizes[*returnSize] = end + 1;
        (*returnSize)++;
        return;
    }
    for (int i = start; i <= end; i++) {
        swap(&nums[start], &nums[i]);
        p(nums, start + 1, end, result, returnSize, columnSizes);
        swap(&nums[start], &nums[i]);
    }
}

int** permute(int* nums, int numsSize, int* returnSize, int** returnColumnSizes) {
    int maxim = 1;
    for (int i = 2; i <= numsSize; i++) {
        maxim *= i;
    }
    int **result = malloc(maxim * sizeof(int*));
    *returnColumnSizes = malloc(maxim * sizeof(int));
    *returnSize = 0;
    p(nums, 0, numsSize - 1, result, returnSize, *returnColumnSizes);
    return result;
}