void nextPermutation(int* nums, int numsSize) {
    int i = numsSize - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) i--;
    if (i >= 0) {
        int j = numsSize - 1;
        while (nums[j] <= nums[i]) j--;
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
    int f = i + 1, s = numsSize - 1;
    while (f < s) {
        int temp = nums[f];
        nums[f] = nums[s];
        nums[s] = temp;
        f++;
        s--;
    }
}