/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* decrypt(int* code, int codeSize, int k, int* returnSize) {
    *returnSize = codeSize;
    int* res = (int*)calloc(codeSize, sizeof(int));
    if(k > 0)
    {
        for(int i = 0; i< codeSize; i++)
        {
            int low = 1;
            int count = i;
            while(low < k+1)
            {
                if(count+low > codeSize-1)
                {
                    count = -low;
                }
                res[i] += code[count+low];
                low++;
            }
        }
    }
    else if(k<0){
        for(int i = 0; i< codeSize; i++)
        {
            int high = -1;
            int count = i;
            while(high > k-1){
                if(count+ high < 0)
                {
                    count = codeSize - high-1;
                }
                res[i] += code[count+high];
                high--;
            }
            //code[i] = j;
        }
    }
    return res;
}