/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* constructTransformedArray(int* nums, int numsSize, int* returnSize) {
    int* res = (int*)malloc(sizeof(int)*numsSize);
    memset(res, 0, sizeof(int) * numsSize);
    *returnSize = numsSize;
    for(int i = 0; i< numsSize; i++)
    {
        res[i] = nums[(i+nums[i]+numsSize*100)%numsSize];
    }
    return res;
}