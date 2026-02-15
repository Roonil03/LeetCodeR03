class Solution {
public:
    int thirdMax(vector<int>& nums) {
        using ll = long long;
        vector<ll> m(3, LLONG_MIN);
        for(int i : nums){
            if (i == m[0] || i == m[1] || i == m[2]){
                continue;
            }
            if (i > m[0]){
                m[2] = m[1];
                m[1] = m[0];
                m[0] = i;
            } else if(i > m[1]){
                m[2] = m[1];
                m[1] = i;
            } else if(i > m[2]){
                m[2] = i;
            }
        }
        return m[2] == LLONG_MIN ? m[0] : m[2];
    }
};