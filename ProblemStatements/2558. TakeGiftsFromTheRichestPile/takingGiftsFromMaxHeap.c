void heapify(int* arr, int n, int i) {
    int largest = i;
    int left = 2* i + 1;
    int right = 2 * i + 2;
    if (left < n && arr[left] > arr[largest])
        largest = left;
    if (right < n && arr[right] > arr[largest])
        largest = right;
    if (largest != i) {
        int temp = arr[i];
        arr[i] = arr[largest];
        arr[largest] = temp;
        heapify(arr, n, largest);
    }
}

void buildMaxHeap(int* arr, int n) {
    for (int i = n / 2 - 1; i >= 0; i--)
        heapify(arr, n, i);
}

long long pickGifts(int* gifts, int giftsSize, int k) {
    buildMaxHeap(gifts, giftsSize);
    for(int i = 0; i< k; i++)
    {
        gifts[0] = (int)pow(gifts[0],0.5);
        heapify(gifts, giftsSize,0);
    }
    long long sum = 0;
    for(int i = 0; i< giftsSize; i++)
    {
        sum += gifts[i];
    }
    return sum;
}