class Solution {
public:
    int intersectionSizeTwo(vector<vector<int>>& intervals) {
        ranges::sort(intervals, [](const auto& a, const auto& b){
            return a[1] != b[1] ? a[1] < b[1] : a[0] > b[0];
        });
        int res {0}, p1 {-1}, p2 {-1};
        for(auto& i : intervals){
            if(i[0] > p2){
                res += 2;
                p1 = i[1] - 1;
                p2 = i[1];
            } else if(i[0] > p1){
                res++;
                p1 = p2;
                p2 = i[1];
            }
        }
        return res;
    }
};