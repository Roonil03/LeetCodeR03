class Solution {
public:
    int partitionString(string s) {
        unordered_set<char> seen;
        int res {1};
        for(char c : s){
            if(seen.count(c)){
                res++;
                seen.clear();
            }
            seen.insert(c);
        }
        return res;
    }
};