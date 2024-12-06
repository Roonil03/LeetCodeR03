int maxCount(int* banned, int bannedSize, int n, int maxSum) {
    long sum = 0;
    int count = 0;
    int hashSet[10001] = {0};
    for(int i = 0; i<bannedSize; i++)
    {
        hashSet[banned[i]] = 1;
    }
    int low = 0, high = n;
    for(int i = 1; i <= n; i++)
    {
        if(i > maxSum || sum >= maxSum || sum+i>maxSum)
        {
            break;
        }
        if(hashSet[i])
        {
            continue;
        }
        sum += i;
        count++;
    }
    return count;
}