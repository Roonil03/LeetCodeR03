class Solution {
public:
    vector<string> topKFrequent(vector<string>& words, int k) {
        unordered_map<string_view, int> count;
        for(auto& i : words){
            count[i]++;
        }
        vector<pair<string_view, int>> c(count.begin(), count.end());
        auto comp = [](const auto& a, const auto& b){
            return a.second != b.second ? a.second > b.second : a.first < b.first;
        };
        if(k < c.size()){
            ranges::nth_element(c, c.begin() + k, (comp));
        }
        ranges::sort(c.begin(), c.begin() + k, comp);
        vector<string> res;
        res.reserve(k);
        for(int i {0}; i < k; i++){
            res.emplace_back(c[i].first);
        }
        return res;
    }
};