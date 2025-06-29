class Solution {
public:
    int hIndex(vector<int>& citations) {
        int n = citations.size();
        int l {0}, r = n - 1, res {0};
        while(l <= r){
            int m = l + (r - l) / 2;
            int h = n - m;
            if(citations[m] >= h){
                res = h;
                r = m - 1;
            } else{
                l = m + 1;
            }
        }
        return res;
    }
};