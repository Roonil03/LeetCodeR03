class Solution {
public:
    int findMaximizedCapital(int k, int w, vector<int>& profits, vector<int>& capital) {
        int n = profits.size();
        vector<pair<int, int>> pj;
        for (int i {0}; i < n; i++){
            pj.emplace_back(capital[i], profits[i]);
        }
        sort(pj.begin(), pj.end());
        priority_queue<int> h;
        int i {0};
        while(k--){
            while(i < n && pj[i].first <= w){
                h.push(pj[i].second);
                i++;
            }
            if (h.empty()){
                break;
            }
            w += h.top();
            h.pop();
        }
        return w;
    }
};