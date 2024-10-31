#include <stdio.h>
#include <stdlib.h>

int minimumMountainRemovals(int* nums, int numsSize) {
    if (numsSize < 3) return numsSize;

    int* leftLIS = (int*)malloc(numsSize * sizeof(int));
    int* rightLIS = (int*)malloc(numsSize * sizeof(int));

    for (int i = 0; i < numsSize; i++) {
        leftLIS[i] = 1;
        rightLIS[i] = 1;
    }
    for (int i = 1; i < numsSize; i++) {
        for (int j = 0; j < i; j++) {
            if (nums[j] < nums[i]) {
                leftLIS[i] = leftLIS[i] > leftLIS[j] + 1 ? leftLIS[i] : leftLIS[j] + 1;
            }
        }
    }
    for (int i = numsSize - 2; i >= 0; i--) {
        for (int j = numsSize - 1; j > i; j--) {
            if (nums[j] < nums[i]) {
                rightLIS[i] = rightLIS[i] > rightLIS[j] + 1 ? rightLIS[i] : rightLIS[j] + 1;
            }
        }
    }
    int maxMountainLength = 0;
    for (int i = 1; i < numsSize - 1; i++) {
        if (leftLIS[i] > 1 && rightLIS[i] > 1) {
            int mountainLength = leftLIS[i] + rightLIS[i] - 1;
            maxMountainLength = maxMountainLength > mountainLength ? maxMountainLength : mountainLength;
        }
    }
    free(leftLIS);
    free(rightLIS);
    return maxMountainLength == 0 ? numsSize : numsSize - maxMountainLength;
}