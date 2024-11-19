long long maximumSubarraySum(int* nums, int numsSize, int k) {
    long long ans = 0, sum = 0;
    int cnt[100001] = {0};
    int l = 0;
    for (int r = 0; r < numsSize; ++r) {
        sum += nums[r];
        cnt[nums[r]]++;        
        while (cnt[nums[r]]>1 || (r-l+1)>k) {
            sum -= nums[l];
            cnt[nums[l]]--;
            l++;
        }        
        if (r - l + 1 == k) {
            ans = (sum > ans) ? sum : ans;
        }
    }    
    return ans;
}