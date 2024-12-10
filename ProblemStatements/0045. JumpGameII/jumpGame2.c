int jump(int* nums, int numsSize) {
    if (numsSize <= 1)
    {
        return 0;
    }
    int count = 0, curr = 0, far = 0;
    for (int i = 0; i < numsSize - 1; i++)
    {
        far = (i + nums[i] > far) ? i + nums[i] : far;
        if (i == curr)
        {
            count++;
            curr = far;
            if (curr >= numsSize - 1)
            {
                break;
            }
        }
    }
    return count;
}