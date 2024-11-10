#include <stdio.h>
#include <limits.h>

#define MAX 32

void update(unsigned** freq, int k, unsigned x, int add) {
    for (int i = 0; i < MAX; i++)
        (*freq)[i] += (add) ? ((x >> i) & 1) : -((x >> i) & 1);
}

int check(unsigned** freq, int k) {
    int b = 0;
    for (int i = 0; i < MAX; i++)
        if (((*freq)[i]) > 0) b |= (1 << i);
    return b >= k;
}

int minimumSubarrayLength(int* nums, int numsSize, int k) {
    unsigned* freq = (unsigned*)calloc(MAX,sizeof(unsigned)*MAX);
    int ans = INT_MAX;
    for (int l = 0, r = 0; r < numsSize; r++) {
        int x = nums[r];
        if (x >= k) return 1;
        update(&freq,k, x, 1);
        while (l < numsSize && check(&freq,k)) {
            if (r - l + 1 < ans) ans = r - l + 1;
            update(&freq, k, nums[l++], 0);
        }
    }
    return (ans == INT_MAX) ? -1 : ans;
}