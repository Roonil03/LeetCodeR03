class Solution {
public:
    int maxProfit(vector<int>& prices, int fee) {
        int n = prices.size();
        if(n < 2){
            return 0;
        }
        int h = -prices[0], res {0};
        for(int i{1}; i < n;i++){
            h = max(h, res- prices[i]);
            res = max(res, h + prices[i] - fee);
        }
        return res;
    }
};