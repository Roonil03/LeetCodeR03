class Solution {
public:
    vector<int> canSeePersonsCount(vector<int>& heights) {
        int n = heights.size();
        vector<int> res(n);
        stack<int> st;
        for(int i = n - 1; i >= 0; i--){
            int c {0};
            while(!st.empty() && st.top() < heights[i]){
                st.pop();
                c++;
            }
            if(!st.empty()){
                c++;
            }
            res[i] = c;
            st.push(heights[i]);
        }
        return res;
    }
};