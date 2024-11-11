#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* searchRange(int* nums, int numsSize, int target, int* returnSize) {
    *returnSize = 2;
    int* res = (int*)malloc(sizeof(int)*2);
    res[0] = -1; res[1] = -1;
    int l = 0, r = numsSize - 1;
    while (l <= r) {
        int mid = l + (r - l) / 2;
        if (nums[mid] == target) {
            res[0] = mid;
            r = mid - 1;
        }
        else if (nums[mid] < target) {
            l = mid + 1;
        }
        else {
            r = mid - 1;
        }
    }
    l = 0;
    r = numsSize - 1;
    while (l <= r) {
        int mid = l + (r - l) / 2;
        if (nums[mid] == target) {
            res[1] = mid;
            l = mid + 1;
            }
        else if (nums[mid] < target) {
            l = mid + 1;
        }
        else {
            r = mid - 1;
        }
    }
    return res;
}