int* resultsArray(int* nums, int numsSize, int k, int* returnSize) {
    *returnSize = numsSize - k + 1;
    int* res = (int*)malloc(sizeof(int) * (*returnSize));
    int isConsecutive = 1;
    for (int i = 1; i < k; i++) {
        if (nums[i] - nums[i - 1] != 1) {
            isConsecutive = 0;
            break;
        }
    }
    res[0] = isConsecutive ? nums[k - 1] : -1;
    for (int i = 1; i < *returnSize; i++) {
        if (isConsecutive && nums[i + k - 1] - nums[i + k - 2] == 1) {
            isConsecutive = 1;
        } else {
            isConsecutive = 1;
            for (int j = i; j < i + k - 1; j++) {
                if (nums[j + 1] - nums[j] != 1) {
                    isConsecutive = 0;
                    break;
                }
            }
        }
        res[i] = isConsecutive ? nums[i + k - 1] : -1;
    }
    return res;
}
