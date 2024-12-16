/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
void heapify(int* arr, int* ids, int n, int i) {
    int smol = i; 
    int l = 2 * i + 1; 
    int r = 2 * i + 2;
    if (l < n && (arr[ids[l]] < arr[ids[smol]] || (arr[ids[l]] == arr[ids[smol]] && ids[l] < ids[smol])))
        smol = l;
    if (r < n && (arr[ids[r]] < arr[ids[smol]] || (arr[ids[r]] == arr[ids[smol]] && ids[r] < ids[smol])))
        smol = r;
    if (smol != i) {
        int temp = ids[i];
        ids[i] = ids[smol];
        ids[smol] = temp;
    heapify(arr, ids, n, smol);
    }
}

int* getFinalState(int* nums, int numsSize, int k, int multiplier, int* returnSize) {
    *returnSize = numsSize;
    int* ids = (int*)malloc(sizeof(int)*numsSize);
    for(int i = 0; i< numsSize; i++)
    {
        ids[i] = i;
    }
    for (int i = numsSize / 2 - 1; i >= 0; i--) {
        heapify(nums, ids, numsSize, i);
    }
    for(int i = 0; i< k; i++)
    {
        nums[ids[0]] *= multiplier;
        heapify(nums, ids, numsSize, 0);
    }
    free(ids);
    return nums;
}