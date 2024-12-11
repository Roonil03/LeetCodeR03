int compare(const void *a, const void *b) {
    return (*(int *)a - *(int *)b);
}

int maximumBeauty(int *nums, int numsSize, int k) {
    qsort(nums, numsSize, sizeof(int), compare);
    int maxBeauty = 0, l = 0;
    for (int r = 0; r < numsSize; r++) {
        while (nums[r] - nums[l] > 2 * k) {
            l++;
        }
        int curr = r - l + 1;
        if (curr > maxBeauty) {
            maxBeauty = curr;
        }
    }
    return maxBeauty;
}