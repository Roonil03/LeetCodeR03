int trap(int* height, int heightSize) {
    if (heightSize <= 2)
        return 0;
    int low = 0, high = heightSize - 1, lmax = 0, hmax = 0, res = 0;
    while (low <= high) {
        if (height[low] < height[high]) {
            if (height[low] >= lmax) 
                lmax = height[low];
            else 
                res += lmax - height[low];
            low++;
        } else {
            if (height[high] >= hmax) 
                hmax = height[high];
            else 
                res += hmax - height[high];
            high--;
        }
    }
    return res;
}
