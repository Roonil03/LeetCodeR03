class Solution {
public:
    int findLucky(vector<int>& arr) {
        unordered_map<int, int> freq;
        for(int n : arr){
            freq[n]++;
        }
        int res {-1};
        for(auto& [num, c] : freq){
            if(num == c){
                res = max(res, num);
            }
        }
        return res;
    }
};