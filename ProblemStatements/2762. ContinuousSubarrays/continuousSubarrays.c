long long continuousSubarrays(int* nums, int numsSize){
    long long count = 0;
    int l = 0;
    int* mini = (int*)malloc(numsSize * sizeof(int));
    int* maxi = (int*)malloc(numsSize * sizeof(int));
    int minFront = 0, minBack = 0, maxFront = 0, maxBack = 0;
    for (int r = 0; r < numsSize; r++)
    {
        while (minBack > minFront && nums[mini[minBack - 1]] >= nums[r]) 
        {
            minBack--;
        }
        while (maxBack > maxFront && nums[maxi[maxBack - 1]] <= nums[r])
        {
            maxBack--;
        }
        mini[minBack++] = r;
        maxi[maxBack++] = r;
        while (nums[maxi[maxFront]] - nums[mini[minFront]] > 2)
        {
            if (mini[minFront] == l) minFront++;
            if (maxi[maxFront] == l) maxFront++;
            l++;
        }
        count += (r - l + 1);
    }
    free(mini);
    free(maxi);
    return count;
}