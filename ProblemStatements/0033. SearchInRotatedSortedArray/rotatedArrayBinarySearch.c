int search(int* nums, int numsSize, int target) {
    int high, low, mid;
    high = numsSize-1;
    low = 0;
    while(low <= high)
    {
        mid = (low + high) / 2;
        if(nums[mid] == target)
        {
            return mid;
        }
        else if(nums[mid] >= nums[low])
        {
            if(nums[low] <= target && target <= nums[mid])
            {
                high = mid- 1;
            }
            else{
                low = mid + 1;
            }
        }
        else{
            if(nums[mid] <= target && target <= nums[high])
                low = mid + 1;
            else
                high = mid - 1;
        }
    }
    return -1;
}