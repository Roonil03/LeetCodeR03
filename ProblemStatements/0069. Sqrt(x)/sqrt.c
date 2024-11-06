#include <stdio.h>

int mySqrt(int x) {
    if (x < 2) return x;
    long left = 2, right = x / 2, mid;
    long sqrt;
    while (left <= right) {
        mid = left + (right - left) / 2;
        sqrt = mid * mid;
        if (sqrt == x) {
            return mid; 
        } else if (sqrt < x) {
            left = mid + 1; 
        } else {
            right = mid - 1; 
        }
    }
    return right;
}
