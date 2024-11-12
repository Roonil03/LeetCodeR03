int searchInsert(int* nums, int numsSize, int target) {
    int low = 0, high = numsSize -1;
    if(target>nums[high]) return numsSize;
    if(target<=nums[low]) return 0;
     while (low <= high) {
        int mid = low + (high - low) / 2;        
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }    
    return low;
}