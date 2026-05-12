class Solution {
public:
    int minimumEffort(vector<vector<int>>& tasks) {
        ranges::sort(tasks, [](const auto&a, const auto& b){
            return (a[1] - a[0]) < (b[1] - b[0]);
        });
        return ranges::fold_left(tasks, 0, [](int r, const auto& i){
            return max(r + i[0], i[1]);
        });
    }
};