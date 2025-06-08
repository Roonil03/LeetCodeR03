class Solution {
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>> res;
        vector<int> c;
        bt(k, n, 1, c, res);
        return res;
    }

    void bt(int k, int n, int s, vector<int>& c, vector<vector<int>>& res) {
        if (n == 0 && c.size() == k) {
            res.push_back(c);
            return;
        }        
        if (c.size() >= k || n <= 0) {
            return;
        }        
        for (int i = s; i <= 9; i++) {
            c.push_back(i);
            bt(k, n - i, i + 1, c, res);
            c.pop_back();
        }
    }
};