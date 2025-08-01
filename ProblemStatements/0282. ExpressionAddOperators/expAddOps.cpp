class Solution {
    using ll = long long;
public:
    // vector<string> addOperators(string num, int target) {
    //     vector<string> res;
    //     dfs(num, target, 0, 0, 0, "", res);
    //     return res;
    // }

    // void dfs(string& nums, int tar, int pos, ll eval, ll md, string path, vector<string>& res){
    //     if (pos == nums.size()){
    //         if(eval == tar){
    //             res.push_back(path);
    //         }
    //         return;
    //     }
    //     for (int i = pos; i < nums.size(); ++i) {
    //         if (i != pos && nums[pos] == '0'){
    //             break;
    //         }
    //         string c = nums.substr(pos, i - pos + 1);
    //         ll cur = stoll(c);
    //         if(pos == 0){
    //             dfs(nums, tar, i + 1, cur, cur, c, res);
    //         } else{
    //             dfs(nums, tar, i + 1, eval + cur, cur, path + "+" + c, res);
    //             dfs(nums, tar, i + 1, eval - cur, -cur, path + "-" + c, res);
    //             dfs(nums, tar, i + 1, eval - md + md * cur, md * cur, path + "*" + c, res);
    //         }
    //     }
    vector<string> ans;
    string s;
    int target;

    vector<string> addOperators(string s, int target) {
        this->s = s;
        this->target = target;
        backtrack( 0, "", 0, 0);
        return ans;
    }

    void backtrack(int i, const string& path, long resSoFar, long prevNum) {
        if (i == s.length()) {
            if (resSoFar == target) ans.push_back(path);
            return;
        }
        string numStr;
        long num = 0;
        for (int j = i; j < s.length(); j++) {
            if (j > i && s[i] == '0'){
                break;
            }
            numStr += s[j];
            num = num * 10 + s[j] - '0';
            if (i == 0) {
                backtrack(j + 1, path + numStr, num, num); 
            } else {
                backtrack(j + 1, path + "+" + numStr, resSoFar + num, num);
                backtrack(j + 1, path + "-" + numStr, resSoFar - num, -num);
                backtrack(j + 1, path + "*" + numStr, resSoFar - prevNum + prevNum * num, prevNum * num);
            }
        }
    }
};