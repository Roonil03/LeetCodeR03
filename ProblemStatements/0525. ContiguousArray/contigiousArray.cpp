class Solution {
    public:
        int findMaxLength(vector<int>& nums) {
            unordered_map<int,int> m;
            int sum = 0;
            unsigned int res = 0;
            for (int i {0}; i < nums.size(); i++){
                sum += (nums[i] == 0)? -1 : 1;
                auto it = m.find(sum);
                if(sum == 0){
                    if (res < i+1){
                        res = i+1;
                    }
                } else if(it != m.end()){
                    if(res < i - it->second){
                        res = i - it->second;
                    }
                } else if(it == m.end()){
                        m.insert({sum, i});
                }
            }
            return res;
        }
    };