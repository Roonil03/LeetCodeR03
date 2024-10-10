int maxWidthRamp(int* nums, int numsSize) {
    int a = 0, b = 0;
    for(int i = 0; i<numsSize;i++)
    {
        //a = i;
        for(int j = i+1;j<numsSize;j++)
        {
            //b = j;
            if(nums[i]<=nums[j] && b-a < j-i)
            {
                a= i;
                b = j;
            }
        }
    }
    return b-a;
}