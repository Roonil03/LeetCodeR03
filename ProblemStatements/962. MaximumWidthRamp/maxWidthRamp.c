int maxWidthRamp(int* nums, int numsSize) {
    int* stack = (int*)malloc(numsSize * sizeof(int));
    int top = -1;
    for (int i = 0; i < numsSize; i++) {
        if (top == -1 || nums[stack[top]] > nums[i]) {
            stack[++top] = i;
        }
    }
    int maxWidth = 0;
    for (int j = numsSize - 1; j >= 0; j--) {
        while (top >= 0 && nums[stack[top]] <= nums[j]) {
            maxWidth = maxWidth > (j - stack[top]) ? maxWidth : (j - stack[top]);
            top--;
        }
    }
    free(stack);
    return maxWidth;
}