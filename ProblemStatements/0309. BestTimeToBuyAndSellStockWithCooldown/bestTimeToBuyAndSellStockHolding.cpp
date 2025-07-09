class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int n = prices.size();
        int h = -prices[0], s {0}, r {0};
        for (int i{1} ; i < n; i++){
            int a = h, b = s, c = r;
            h = max(a, c - prices[i]);
            s = a + prices[i];
            r = max(c, b);
        }
        return max(s, r);
    }
};