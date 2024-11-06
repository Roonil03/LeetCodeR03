#include <stdio.h>
#include <stdbool.h>

int setBits(int num) {
    int count = 0;
    for (int i = 31; i >= 0; i--) {
        if ((num >> i) & 1) {
            count++;
        }
    }
    return count;
}

bool check(int* nums, int numsSize) {
    for (int i = 0; i < numsSize - 1; i++) {
        if (nums[i] > nums[i + 1]) {
            return false;
        }
    }
    return true;
}

bool canSortArray(int* nums, int numsSize) {
    int* count = (int*)malloc(numsSize * sizeof(int));
    for (int i = 0; i < numsSize; i++) {
        count[i] = setBits(nums[i]);
    }
    int k = 0;
    while (k < numsSize) {
        for (int i = 1; i < numsSize; i++) {
            if (count[i] == count[i - 1] && nums[i] < nums[i - 1]) {
                int temp = nums[i];
                nums[i] = nums[i - 1];
                nums[i - 1] = temp;
            }
        }
        if (check(nums, numsSize)) {
            free(count);
            return true;
        }
        k++;
    }
    free(count);
    return false;
}