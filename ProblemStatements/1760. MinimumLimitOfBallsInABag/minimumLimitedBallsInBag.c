# include <stdbool.h>

bool canDistribute(int *nums, int numsSize, int maxOperations, int maxP) {
    int ops = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] > maxP) {
            ops += (nums[i] - 1) / maxP;
            if (ops > maxOperations) {
                return false;
            }
        }
    }
    return true;
}

int minimumSize(int* nums, int numsSize, int maxOperations) {
    int low = 1, high = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] > high) {
            high = nums[i];
        }
    }
    while (low < high) {
        int mid = low + (high - low) / 2;
        if (canDistribute(nums, numsSize, maxOperations, mid)) {
            high = mid;
        } else {
            low = mid + 1;
        }
    }
    return low;
}