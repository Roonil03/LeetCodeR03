class Solution {
public:
    int maxDistance(vector<int>& colors) {
        int n = colors.size();
        int l = {0}, r {0};
        for(int i {0}; i < n; i++){
            if ((colors[i] ^ colors.back()) | (colors[n - 1 - i] ^ colors[0])){
                return n - 1 - i;
            }
        }
        return 0;
    }
};