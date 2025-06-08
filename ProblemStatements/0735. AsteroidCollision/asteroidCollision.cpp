class Solution {
public:
    vector<int> asteroidCollision(vector<int>& asteroids) {
        vector<int> st;        
        for(int ast : asteroids) {
            while(!st.empty() && ast < 0 && st.back() > 0) {
                int sum = ast + st.back();
                if(sum < 0) {
                    st.pop_back();
                    continue;
                }
                if(sum > 0) {
                    ast = 0;
                    break;
                }
                if(sum == 0) {
                    st.pop_back();
                    ast = 0;
                    break;
                }
            }
            if(ast != 0) {
                st.push_back(ast);
            }
        }        
        return st;
    }
};