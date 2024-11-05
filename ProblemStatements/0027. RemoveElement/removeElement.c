void sort(int* nums, int start, int end) {
    for (int i = start; i < end - 1; i++) {
        for (int j = start; j < end - i - 1; j++) {
            if (nums[j] > nums[j + 1]) {
                int temp = nums[j];
                nums[j] = nums[j + 1];
                nums[j + 1] = temp;
            }
        }
    }
}
int removeElement(int* nums, int numsSize, int val) {
    int j = 0;
    for (int i = 0; i < numsSize; i++) {
        if (nums[i] != val) {
            nums[j] = nums[i];
            j++;
        }
    }
    sort(nums,0,j);
    return j;
}
