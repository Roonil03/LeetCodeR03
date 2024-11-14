#include <stdbool.h>

bool canDistribute(int* quantities, int quantitiesSize, int n, int k) {
    int count = 0;
    for (int i = 0; i < quantitiesSize; i++) {
        count += (quantities[i] + k - 1) / k;
        if (count > n) {
            return false;
        }
    }
    return true;
}

int minimizedMaximum(int n, int* quantities, int quantitiesSize) {
    int low = 1, high = 0;
    for (int i = 0; i < quantitiesSize; i++) {
        if (quantities[i] > high) {
            high = quantities[i];
        }
    }
    while (low < high) {
        int mid = low + (high - low) / 2;
        if (canDistribute(quantities, quantitiesSize, n, mid)) {
            high = mid;
        } else {
            low = mid + 1;
        }
    }
    return low;
}
