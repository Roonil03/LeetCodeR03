#include <stdbool.h>

typedef struct {
    int val;
    int id;
} Data;

int compare(const void* a, const void* b) {
    Data* da = (Data*)a;
    Data* db = (Data*)b;
    return da->val - db->val;
}

long long findScore(int* nums, int numsSize) {
    bool* marked = (bool*)calloc(numsSize, sizeof(bool));
    long long score = 0;
    Data* heap = (Data*)malloc(numsSize * sizeof(Data));
    for (int i = 0; i < numsSize; i++) {
        heap[i].val = nums[i];
        heap[i].id = i;
    }
    qsort(heap, numsSize, sizeof(Data), compare);
    for (int i = 0; i < numsSize; i++) {
        int id = heap[i].id;
        if (marked[id]) continue;
        score += nums[id];
        marked[id] = true;
        if (id > 0 && !marked[id - 1]) {
            marked[id - 1] = true;
        }
        if (id < numsSize - 1 && !marked[id + 1]) {
            marked[id + 1] = true;
        }
    }
    free(marked);
    free(heap);
    return score;
}
