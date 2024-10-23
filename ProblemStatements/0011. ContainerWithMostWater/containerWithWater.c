#include <stdio.h>
#include <stdlib.h>

int mini(int a, int b) {
    return a < b ? a : b;
}
int maxArea(int* height, int heightSize) {
    int left = 0;
    int right = heightSize - 1;
    int maxArea = 0;

    while (left < right) {
        int currentArea = mini(height[left], height[right]) * (right - left);
        if (currentArea > maxArea) {
            maxArea = currentArea;
        }
        if (height[left] < height[right]) {
            left++;
        } else {
            right--;
        }
    }
    return maxArea;
}
// int main() {
//     int height[] = {4, 3, 2, 1, 4};
//     int heightSize = sizeof(height) / sizeof(height[0]);
   
//     int result = maxArea(height, heightSize);
//     printf("Maximum area of water that can be contained: %d\n", result);
   
//     return 0;
// }