#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef struct{
    int * arr;
    int size;
    int capacity;
}Heap;

Heap* createHeap(int capacity)
{
    Heap* heap = (Heap*)malloc(sizeof(Heap));
    heap->capacity = capacity;
    heap->size = 0;
    heap->arr = (int*)malloc(capacity*sizeof(int));
    return heap;
}

void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void heapify(Heap* heap, int index)
{
    int largest = index;
    int left = 2*index + 1;
    int right = 2* index + 2;
    if(left <heap->size && heap ->arr[left] > heap->arr[largest])
    {
        largest = left;
    }
    if(right <heap->size && heap->arr[right] > heap->arr[largest])
    {
        largest = right;
    }
    if(largest != index)
    {
        swap(&heap->arr[index],&heap->arr[largest]);
        heapify(heap, largest);
    }
}
void insert(Heap* heap, int val)
{
    if(heap->size == heap->capacity)
    {
        printf("Heap is full\n");
        return;
    }
    heap->arr[heap->size] = val;
    heap->size++;
    for(int i = heap->size / 2 - 1; i>= 0; i--)
    {
        heapify(heap, i);
    }
}

int replaceMax(Heap* heap)
{
    if(heap->size <= 0)
    {
        printf("Heap is empty!\n");
        return -1;
    }
    if(heap->size == 1)
    {
        int n = heap->arr[0];
        heap->arr[0] = (int)ceil(heap->arr[0]/3.0);
        //heap->size--;
        return n;
    }
    int root = heap->arr[0];
    heap->arr[0] = (int)ceil(heap->arr[0]/3.0);
    heapify(heap,0);
    return root;
}

void freeHeap(Heap* heap) {
    free(heap->arr);
    free(heap);
}

long long maxKelements(int* nums, int numsSize, int k) {
    Heap* heap = createHeap(numsSize);
    for(int i = 0; i< numsSize;i++)
    {
        insert(heap, nums[i]);
    }
    long long sum = 0;
    for(int j = 0; j<k;j++)
    {
        sum += replaceMax(heap);
    }
    freeHeap(heap);
    return sum;
}
/*int main() {
    int nums[] = {1, 10, 3, 3, 3};
    int k = 3;
    long long result = maxKelements(nums, 5, k);
    printf("The output is: %lld\n", result);
    return 0;
}*/