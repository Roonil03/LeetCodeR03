#include <stdio.h>

int backtrack(int* nums, int numsSize, int index, int currentOr, int maxOr) {
    if (index == numsSize) {
        return currentOr == maxOr ? 1 : 0;
    }
    int include = backtrack(nums, numsSize, index + 1, currentOr | nums[index], maxOr);
    int exclude = backtrack(nums, numsSize, index + 1, currentOr, maxOr);
    return include + exclude;
}

int countMaxOrSubsets(int* nums, int numsSize) {
    int maxOr = 0;
    for (int i = 0; i < numsSize; i++) {
        maxOr |= nums[i];
    }
    return backtrack(nums, numsSize, 0, 0, maxOr);
}
