class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        int n = temperatures.size();
        vector<int>res(n, 0);
        stack<int> st;
        for(int i{0}; i < n; i++){
            while(!st.empty() && temperatures[i] > temperatures[st.top()]){
                int id = st.top();
                st.pop();
                res[id] = i - id;
            }
            st.push(i);
        }
        return res;
    }
};