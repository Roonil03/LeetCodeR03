class Solution {
public:
    using ll = long long;

    bool find132pattern(vector<int>& nums) {
        int n = nums.size();
        if (n < 3){
            return false;
        }
        stack<ll> st;
        ll m = numeric_limits<ll>::min();
        for(int i = n - 1; i >= 0; i--){
            if(nums[i] < m){
                return true;
            }
            while(!st.empty() && st.top() < nums[i]){
                m = max(m, st.top());
                st.pop();
            }
            st.push(nums[i]);
        }
        return false;
    }
};
