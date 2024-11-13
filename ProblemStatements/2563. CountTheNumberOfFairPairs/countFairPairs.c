int lowerBound(int* arr, int size, int start, int target) {
    int left = start, right = size - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (arr[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;
}

int upperBound(int* arr, int size, int start, int target) {
    int left = start, right = size - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (arr[mid] <= target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return left;
}

int compare(const void *a, const void *b) {
    return (*(int *)a - *(int *)b);
}

long long countFairPairs(int* nums, int numsSize, int lower, int upper) {
    qsort(nums, numsSize, sizeof(int), (int (*)(const void *, const void *))&compare);    
    long long count = 0;
    for (int i = 0; i < numsSize; i++) {
        int lowIndex = lowerBound(nums, numsSize, i + 1, lower - nums[i]);
        int highIndex = upperBound(nums, numsSize, i + 1, upper - nums[i]);
        count += highIndex - lowIndex;
    }
    return count;
}
