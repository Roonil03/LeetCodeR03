/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdlib.h>

int* getMaximumXor(int* nums, int numsSize, int maximumBit, int* returnSize) {
    int all_bits_set = (1 << maximumBit) - 1;
    int prefixXOR = 0;
    for (int i = 0; i < numsSize; i++) {
        prefixXOR ^= nums[i];
    }
    int* answer = (int*)malloc(numsSize * sizeof(int));
    *returnSize = numsSize;
    for (int i = 0; i < numsSize; i++) {
        answer[i] = prefixXOR ^ all_bits_set;
        prefixXOR ^= nums[numsSize - 1 - i];  
    }

    return answer;
}
