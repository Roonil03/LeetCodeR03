class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        int count {0};
        for (int i {0}; i < flowerbed.size() && count < n; i++){
            if (flowerbed[i] == 0){
                int n = (i == flowerbed.size() - 1) ? 0 : flowerbed[i + 1];
                int p = (i == 0) ? 0 : flowerbed[i-1];
                if(n == 0 && p == 0){
                    flowerbed[i] = 1;
                    count++;
                }
            }
        }
        return count == n;
    }
};